package md5

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"testing"
	"testing/quick"

	"github.com/ReconfigureIO/crypto/md5/host"
)

func Sum(bs []byte) []byte {
	b := host.Pad(bs)
	buff := bytes.NewBuffer(b)

	arr := [16]uint32{}

	dig := New()

	i := len(b) / 4
	for i != 0 {
		binary.Read(buff, binary.LittleEndian, &arr)
		dig = dig.Block(arr)
		i -= 16
	}
	ret := dig.Sum()
	return ret[:]
}

func ShouldBeReference(t *testing.T, message string, reference string) {
	b := Sum([]byte(message))

	s := hex.EncodeToString(b)
	if s != reference {
		t.Errorf("Expected reference \"%s\" to be \"%s\" got \"%s\"", message, reference, s)
	}
}

func TestReference(t *testing.T) {
	ShouldBeReference(t, "", "d41d8cd98f00b204e9800998ecf8427e")
	ShouldBeReference(t, "The quick brown fox jumps over the lazy dog", "9e107d9d372bb6826bd81d3542a419d6")
	ShouldBeReference(t, "The quick brown fox jumps over the lazy dog.", "e4d909c290d0fb1ca068ffaddf22cbd0")
}

func TestMain(t *testing.T) {
	ours := func(bytes []byte) string {
		b := Sum(bytes)
		return hex.EncodeToString(b)
	}

	theirs := func(bytes []byte) string {
		b := md5.Sum(bytes)
		return hex.EncodeToString(b[:])
	}

	if err := quick.CheckEqual(ours, theirs, nil); err != nil {
		t.Error(err)
	}
}
