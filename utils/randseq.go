package utils

import (
	"math/rand"
	"time"
)

const (
	NUMS             = "0123456789"
	LETTERS          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LETTERS_AND_NUMS = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandSeq(sampleSet string, seqLen int) string {
	b := make([]byte, seqLen)
	for i := range b {
		b[i] = sampleSet[rand.Intn(len(sampleSet))]
	}
	return string(b)
}

func RandNumSeq(seqLen int) string {
	return RandSeq(NUMS, seqLen)
}

func RandLetterSeq(seqLen int) string {
	return RandSeq(LETTERS, seqLen)
}

func RandLetterAndNumSeq(seqLen int) string {
	return RandSeq(LETTERS_AND_NUMS, seqLen)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
