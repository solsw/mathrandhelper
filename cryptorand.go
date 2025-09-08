package mathrandhelper

import (
	"bufio"
	cryptorand "crypto/rand"
	"encoding/binary"
	"io"
	mathrandv2 "math/rand/v2"
)

// NewCryptoRand returns [math/rand/v2/Rand] instance backed by [crypto/rand].
func NewCryptoRand() *mathrandv2.Rand {
	crs := cryptoRandSource{byteReader: bufio.NewReader(cryptorand.Reader)}
	return mathrandv2.New(crs)
}

// cryptoRandSource implements the [math/rand/v2/Source] interface.
type cryptoRandSource struct {
	// inspired by github.com/andrew-d/csmrand
	byteReader io.ByteReader
}

// Uint64 implements the [math/rand/v2/Source.Uint64] method.
func (s cryptoRandSource) Uint64() uint64 {
	u64, _ := binary.ReadUvarint(s.byteReader)
	return u64
}
