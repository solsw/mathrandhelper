package mathrandhelper

import (
	"bufio"
	cryptorand "crypto/rand"
	"encoding/binary"
	"io"
	mathrand "math/rand"
)

// NewCryptoRand returns [math/rand.Rand] instance backed by [crypto/rand].
func NewCryptoRand() *mathrand.Rand {
	crs := cryptoRandSource{bufio.NewReader(cryptorand.Reader)}
	return mathrand.New(crs)
}

// cryptoRandSource implements the [math/rand.Source] interface.
type cryptoRandSource struct {
	// inspired by github.com/andrew-d/csmrand
	byteReader io.ByteReader
}

// Int63 implements the [math/rand.Source.Int63] method.
func (s cryptoRandSource) Int63() int64 {
	u64, _ := binary.ReadUvarint(s.byteReader)
	return int64(u64 >> 1)
}

// Seed implements the [math/rand.Source.Seed] method.
func (cryptoRandSource) Seed(int64) {}
