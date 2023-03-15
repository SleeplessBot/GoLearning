package examples

import (
	"fmt"
	"golearning/utils"
	"time"
)

func ExampleKvs() {
	kvs, _ := utils.NewKVS()

	fmt.Println(kvs.Get("abc"))
	fmt.Println(kvs.Set("abc", "test"))
	fmt.Println(kvs.Get("abc"))
	fmt.Println(kvs.Del("abc"))
	fmt.Println(kvs.Get("abc"))

	ts := time.Now().UnixMilli()
	for i := 0; i < 1000000; i++ {
		kvs.Set("abc", "test")
	}
	te := time.Now().UnixMilli()
	fmt.Println(te - ts)

	ts = time.Now().UnixMilli()
	for i := 0; i < 1000000; i++ {
		kvs.Get("abc")
	}
	te = time.Now().UnixMilli()
	fmt.Println(te - ts)
}
