package examples

import (
	"fmt"
	"golearning/utils"
	"os"
	"strings"
)

func ExampleCas() {
	cas := utils.DefaultCas

	fmt.Println("BaseDir:")
	fmt.Println(cas.BaseDir())

	sr := strings.NewReader("abcdefg")
	key, err := cas.Create(sr)
	fmt.Println("Create:")
	fmt.Println(key, err)

	filePath, err := cas.GetFilePath(key)
	fmt.Println("GetFilePath - T:")
	fmt.Println(filePath, err)
	fmt.Println("GetFilePath - F:")
	fmt.Println(cas.GetFilePath("12345"))

	key, err = cas.FilePathToKey(filePath)
	fmt.Println("FilePathToKey - T:")
	fmt.Println(key, err)
	fmt.Println("FilePathToKey - F:")
	fmt.Println(cas.FilePathToKey("abc/def"))

	fmt.Println("IsExist - T:")
	fmt.Println(cas.IsExist(key))
	fmt.Println("IsExist - F:")
	fmt.Println(cas.IsExist("12345"))

	tmpFileName := "tmp_file_000001.txt"
	tmpFile, _ := os.Create(tmpFileName)
	tmpFile.Write([]byte("1234567890"))
	tmpFile.Close()
	filePath, key, err = cas.MoveFileToCas(tmpFileName)
	fmt.Println("MoveFileToCas:")
	fmt.Println(filePath, key, err)

	fmt.Println("Delete - T:")
	fmt.Println(cas.Delete(key))
	fmt.Println("Delete - F:")
	fmt.Println(cas.Delete("1-2-3-4"))
}
