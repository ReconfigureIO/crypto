package host

import (
	"bytes"
	"encoding/binary"
)

func Pad(b []byte) []byte {
	len := uint64(len(b))

	var buff bytes.Buffer
	// add 1 bit
	buff.Write([]byte{0x80})

	// append "0" bit until message length in bits ≡ 448 (mod 512)
	remaining := (len + 1) % 64
	toWrite := 56 - remaining

	if remaining >= 56 {
		toWrite = 64 + toWrite
	}

	// pad with zeros
	buff.Write(make([]byte, toWrite))

	// write length in bits
	binary.Write(&buff, binary.LittleEndian, len<<3)

	return append(b, buff.Bytes()...)

}
