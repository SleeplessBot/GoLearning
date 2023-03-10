package experiments

import (
	"fmt"
	"reflect"
)

func PrintObjectType(object interface{}) {
	fmt.Printf("type: %s\n", reflect.TypeOf(object))
}

func ExperimentStringElementType() {
	s := "hello世界"
	PrintObjectType(s)    // string
	PrintObjectType(s[1]) // uint8, byte is an alias of uint8
	for i, v := range s {
		PrintObjectType(i) // int
		PrintObjectType(v) // int32, rune is an alias of int32, a rune is a unicode char
		break
	}
	fmt.Printf("string len:%d\n", len(s))       // 11
	fmt.Printf("rune len:%d\n", len([]rune(s))) // 7
}
