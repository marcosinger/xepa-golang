package main

import "fmt"

// start func OMIT
func main() {
	message := Hello("Rocky Balboa")
	fmt.Println(message)
}

func Hello(name string) string { // HL
	return fmt.Sprintf("Hey %s, my friend..", name)
}

// end func OMIT
