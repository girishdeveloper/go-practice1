package main

import (
	"fmt"
)

func main() {
	fmt.Println("Behind the scene, a string is a collection of bytes.")
	fmt.Println("We can iterate over the bytes of a string with a for loop")
	s := "എനിക്ക് Golang ഇഷ്ടമാണ്"
	fmt.Printf("iterate on %s\n", s)
	for _, v := range s {
		fmt.Println(v) // this will give you decimal values instead on the character
		fmt.Printf("Unicode point %U - character %c - binary %b - hex %X - decimal %d\n", v, v, v, v, v)
	}
	fmt.Println("runes can be created by using single quotes.")
	var Rune rune = 'Z'
	fmt.Printf("A rune is %c = %U, %d\n", Rune, Rune, Rune)
}
