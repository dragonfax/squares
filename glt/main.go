package glt

import (
	"fmt"
)

func test(c coreWidget) bool {
	return true
}

func main() {
	fmt.Println(test(&Center{}))
}
