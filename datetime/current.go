package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("print the current date and time")
	fmt.Printf("current date and time is: %s\n", start)
	fmt.Printf("time since is %s\n", time.Since(start))
	fmt.Printf("time until is %s\n", time.Until(start))
	fmt.Printf("time elapsed is %s\n", time.Now().Sub(start))
}
