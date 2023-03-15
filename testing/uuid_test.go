package testing

import (
	"golearning/utils"
	"testing"
)

func BenchmarkB62UUID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.Base62UUID()
	}
}

func BenchmarkB16UUID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		utils.Base16UUID()
	}
}
