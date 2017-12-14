package main

import (
	// import the entire framework (including bundled verilog)
	_ "github.com/ReconfigureIO/sdaccel"
	// Use the new AXI protocol package
	aximemory "github.com/ReconfigureIO/sdaccel/axi/memory"
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"

	"github.com/ReconfigureIO/crypto/md5"
)

// Read the number of blocks from
func ProcessMD5(
	numBlocks uint,
	inputData uintptr,
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData) md5.Digest {

	d := md5.New()

	blocks := make(chan [16]uint32)

	go func() {
		num32s := uint32(numBlocks << 4)
		block := [16]uint32{}

		data := make(chan uint32)

		go aximemory.ReadBurstUInt32(
			memReadAddr, memReadData, true, inputData, num32s, data)

		for i := numBlocks; i != 0; i-- {
			for j := 0; j != 16; j += 1 {
				block[j] = <-data
			}
			blocks <- block
		}
	}()

	for i := numBlocks; i != 0; i-- {
		d = d.Block(<-blocks)
	}

	return d
}

func WriteSum(
	d md5.Digest,
	outputData uintptr,
	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	vals := d.Sum()
	data := make(chan uint8)
	go func() {
		for i := 0; i < 16; i++ {
			data <- uint8(vals[i])
		}
	}()

	aximemory.WriteBurstUInt8(
		memWriteAddr, memWriteData, memWriteResp, true, outputData, 16, data)
}

func Top(
	// The number of blocks to process
	numBlocks uint,

	inputData uintptr,
	outputData uintptr,

	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	d := ProcessMD5(numBlocks, inputData, memReadAddr, memReadData)
	WriteSum(d, outputData, memWriteAddr, memWriteData, memWriteResp)
}
