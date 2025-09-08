package mathrandhelper

import (
	"encoding/binary"
	"math/rand/v2"
)

// RandomBuf fills 'buf' with random bytes.
func RandomBuf(buf []byte) {
	if len(buf) == 0 {
		return
	}
	bbu64 := make([]byte, 8)
	for i := range len(buf) / 8 {
		binary.NativeEndian.PutUint64(bbu64, rand.Uint64())
		copy(buf[i*8:], bbu64)
	}
	remainder := len(buf) % 8
	if remainder == 0 {
		return
	}
	binary.NativeEndian.PutUint64(bbu64, rand.Uint64())
	copy(buf[len(buf)/8*8:], bbu64[0:remainder])
}
