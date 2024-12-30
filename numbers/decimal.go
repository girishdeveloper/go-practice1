package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("decimal value: %T %d\n", int64(193), strconv.FormatInt(193, 10))
	fmt.Printf("binary value: %T %b\n", 123, strconv.FormatInt(123, 2))
	//fmt.Printf("boolean value: %T %v\n", true, true)
	fmt.Printf("octal value: %T %o\n", 0722, strconv.FormatInt(0722, 8))
	fmt.Printf("hexadecimal value: %T %v\n", 0x21a4, strconv.FormatInt(8612, 16))
}
