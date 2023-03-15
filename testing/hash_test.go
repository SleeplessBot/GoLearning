package testing

import (
	"golearning/utils"
	"testing"
)

func BenchmarkMd5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.GetFileMd5("main.go")
	}
}

func BenchmarkSha1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.GetFileSha1("main.go")
	}
}
