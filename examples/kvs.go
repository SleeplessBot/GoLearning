package examples

import (
	"fmt"
	"golearning/utils"
	"time"
)

func ExampleKvs() {
	fmt.Println(utils.DefaultKvs.Get("abc"))
	fmt.Println(utils.DefaultKvs.Set("abc", "test"))
	fmt.Println(utils.DefaultKvs.Get("abc"))
	fmt.Println(utils.DefaultKvs.Del("abc"))
	fmt.Println(utils.DefaultKvs.Get("abc"))

	ts := time.Now().UnixMilli()
	for i := 0; i < 1000000; i++ {
		utils.DefaultKvs.Set("abc", "test")
	}
	te := time.Now().UnixMilli()
	fmt.Println(te - ts)

	ts = time.Now().UnixMilli()
	for i := 0; i < 1000000; i++ {
		utils.DefaultKvs.Get("abc")
	}
	te = time.Now().UnixMilli()
	fmt.Println(te - ts)
}
