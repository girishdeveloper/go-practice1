package main

import (
	"fmt"
)

func main() {
	fmt.Println("In Go language, strings are immutable")
	fmt.Println(" i.e. they cannot be modified once created.")
	fmt.Println("Types of strings:")
	fmt.Println("01. raw strings: These are defined between back quotes.")
	fmt.Println("\t\tForbidden character: back quotes")
	fmt.Println("\t\tDiscarded character: carriage returns \\r")
	fmt.Println("02. interpreted strings: These are defined between double quotes")
	fmt.Println("\t\tForbidden characters: \\n and unescaped double quotes")
	fmt.Println("Lets see the examples.")
	var raw = `raw string follows:
	browsing under an umbrella 
	at the picture-book store.`
	fmt.Println(raw)
	var interpreted = "interpreted: i love spring"
	fmt.Println(interpreted)
}
