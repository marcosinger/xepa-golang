package main

import "fmt"

// start func OMIT
func main() {
  name := Hello("Rocky Balboa")
  fmt.Println(name);
}

func Hello(name string) string { // HL
  return fmt.Sprintf("Hey %s, my friend..", name)
}
// end func OMIT
