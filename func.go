package main

import "fmt"

// start func OMIT
func main() {
	// Existem 3 formas de declarar uma variável:

	// var message string
	// message = Hello("Rocky Balboa")

	// var message = Hello("Rocky Balboa")

	// essa seria uma sugar syntax (somente dentro de funções)
	message := Hello("Rocky Balboa")
	fmt.Println(message)
}

func Hello(name string) string { // HL
	return fmt.Sprintf("Hey %s, my friend..", name)
}

// end func OMIT
