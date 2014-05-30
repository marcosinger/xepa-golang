package main

import "fmt"

type Fighter struct {
	Name string
	Age  int
}

// start interface OMIT
func (f Fighter) String() string { // HL
	return fmt.Sprintf("Ele vem ai! %s/%d", f.Name, f.Age)
}

// end interface OMIT

func main() {
	// start print interface OMIT
	rocky := Fighter{Age: 25, Name: "Rocky Balboa"}
	fmt.Println(rocky)
	// end print interface OMIT
}
