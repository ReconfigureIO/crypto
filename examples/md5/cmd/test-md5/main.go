package main

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/ReconfigureIO/sdaccel/xcl"
	"log"

	"github.com/ReconfigureIO/crypto/md5/host"
)

func main() {
	world := xcl.NewWorld()
	defer world.Release()

	krnl := world.Import("kernel_test").GetKernel("reconfigure_io_sdaccel_builder_stub_0_1")
	defer krnl.Release()

	msg := host.Pad([]byte("The quick brown fox jumps over the lazy dog"))
	msgSize := binary.Size(msg)

	inputBuff := world.Malloc(xcl.ReadOnly, uint(msgSize))
	defer inputBuff.Free()

	outputBuff := world.Malloc(xcl.ReadOnly, 16)
	defer outputBuff.Free()

	binary.Write(inputBuff.Writer(), binary.LittleEndian, msg)
	numBlocks := uint32(msgSize / 64)

	krnl.SetArg(0, numBlocks)
	krnl.SetMemoryArg(1, inputBuff)
	krnl.SetMemoryArg(2, outputBuff)

	krnl.Run(1, 1, 1)

	ret := make([]byte, 16)
	err := binary.Read(outputBuff.Reader(), binary.LittleEndian, ret)
	if err != nil {
		log.Fatal("binary.Read failed:", err)
	}

	s := hex.EncodeToString(ret)
	log.Printf("Got hex string of %s", s)

	if s != "9e107d9d372bb6826bd81d3542a419d6" {
		log.Fatalf("%s != %s", s, "9e107d9d372bb6826bd81d3542a419d6")
	}
}
