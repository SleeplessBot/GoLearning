package examples

import (
	"fmt"
	"golearning/utils"
)

func ExampleBase62UUID() {
	fmt.Printf("base62 uuid: %s\n", utils.Base62UUID())
}
