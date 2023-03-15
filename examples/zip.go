package examples

import (
	"fmt"
	"golearning/utils"
)

func ExampleZipAndUnzip() {
	fmt.Println(utils.SaveStrToFile(".tmp/ab/a.txt", "aaaaa"))
	fmt.Println(utils.SaveStrToFile(".tmp/ab/b.txt", "bbbbb"))
	fmt.Println(utils.CreateZipFromDir(".tmp/ab", ".tmp/ab.zip"))

	fmt.Println(utils.SaveStrToFile(".tmp/c/c.txt", "aaaaa"))
	fmt.Println(utils.SaveStrToFile(".tmp/d/d.txt", "bbbbb"))
	fmt.Println(utils.CreateZipFromFiles([]string{".tmp/c/c.txt", ".tmp/d/d.txt"}, ".tmp/cd.zip", "cd"))

	fmt.Println(utils.Unzip(".tmp/ab.zip", ".tmp/ab_unzip"))
	fmt.Println(utils.Unzip(".tmp/cd.zip", ".tmp/cd_unzip"))
}
