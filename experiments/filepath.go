package experiments

import (
	"fmt"
	"path/filepath"
)

func ExperimentFilePath() {
	fmt.Println(filepath.Join("a", "b"))     // a\b
	fmt.Println(filepath.Join("a", "./b"))   // a\b
	fmt.Println(filepath.Join("./a", "./b")) // a\b
	fmt.Println(filepath.Join("", "a"))      // a
	fmt.Println(filepath.Join("b", ""))      // b
	fmt.Println(filepath.Join("a/b", "c"))   // a\b\c
	fmt.Println(filepath.Join("a", "b/c"))   // a\b\c
	fmt.Println(filepath.Join("a/b", "c/d")) // a\b\c\d

	fmt.Println(filepath.Base("a/b"))  // b
	fmt.Println(filepath.Base("a/b/")) // b
	fmt.Println(filepath.Base("a"))    // a
	fmt.Println(filepath.Base("a/"))   // a
	fmt.Println(filepath.Base("."))    // .

	fmt.Println(filepath.Abs("main.go")) // D:\github\SleeplessBot\GoLearning\main.go

	fmt.Println(filepath.Clean("a/b\\c"))   // a\b\c
	fmt.Println(filepath.Clean("a/b/../c")) // a\c

	fmt.Println(filepath.Dir("a/b"))  // a
	fmt.Println(filepath.Dir("a/b/")) // a\b
	fmt.Println(filepath.Dir("a/"))   // a
	fmt.Println(filepath.Dir("a"))    // .

	fmt.Println(filepath.Ext("a/b.txt"))     // .txt
	fmt.Println(filepath.Ext("a/b.txt.jpg")) // .jpg
	fmt.Println(filepath.Ext("a/b"))         //

	fmt.Println(filepath.FromSlash("a/b/c")) //a\b\c
}
