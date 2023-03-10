package utils

import (
	"encoding/hex"
	"math/big"

	"github.com/google/uuid"
)

func Base62UUID() string {
	var i big.Int
	newUuid := uuid.New()
	i.SetBytes(newUuid[:])
	return i.Text(62)
}

func Base16UUID() string {
	var buf [32]byte
	newUuid := uuid.New()
	hex.Encode(buf[:], newUuid[:])
	return string(buf[:])
}
