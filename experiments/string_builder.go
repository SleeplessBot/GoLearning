package experiments

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func ExperimentStringsBuilder() {
	st := time.Now().UnixMilli()
	var buffer bytes.Buffer
	for i := 0; i < 1000000; i++ {
		buffer.WriteString("a")
	}
	et := time.Now().UnixMilli()
	fmt.Println(et - st) // 4

	st = time.Now().UnixMilli()
	var sb strings.Builder
	for i := 0; i < 1000000; i++ {
		sb.WriteString("a")
	}
	et = time.Now().UnixMilli()
	fmt.Println(et - st) // 2

	st = time.Now().UnixMilli()
	var s string
	for i := 0; i < 1000000; i++ {
		s += "a"
	}
	et = time.Now().UnixMilli()
	fmt.Println(et - st) // 53110
}
