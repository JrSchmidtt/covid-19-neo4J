package id

import (
	"fmt"
	"time"

	"github.com/oklog/ulid"
	"github.com/richardlehane/crock32"
)

var idRng LockedPcg128
var keyRng LockedPcg128

func init() {
	t := uint64(time.Now().UnixNano())
	idRng = NewLockedPcg128()
	idRng.Seed(t, t)
	a, b := idRng.Uint128()
	keyRng = NewLockedPcg128()
	keyRng.Seed(a, b)
	err := crock32.SetDigits("0123456789ABCDEFGHJKMNPQRSTVWXYZ"); if err != nil {
		fmt.Println("Error setting crock32 digits")
	}
}

func NewRequestId() string {
	return NewRequestIdNano(time.Now().UTC().UnixNano())
}

func NewRequestIdNano(nano int64) string {
	t := uint64(nano / int64(time.Millisecond))
	x, y := idRng.Uint128()
	return ulid.ULID{
		byte(t >> 40), byte(t >> 32), byte(t >> 24), byte(t >> 16),
		byte(t >> 8), byte(t), byte(y >> 56), byte(y >> 48),
		byte(y >> 40), byte(y >> 32), byte(y >> 24), byte(y >> 16),
		byte(y >> 8), byte(y), byte(x >> 56), byte(x >> 48),
	}.String()
}

// New random ID.
func NewId() string {
	return crock32.Encode(idRng.Uint64())
}

// Create a new ID with a prefix.
// The prefix is added to the beginning of the ID.
func NewIdWithPrefix(prefix string) string {
	return prefix + NewId()
}