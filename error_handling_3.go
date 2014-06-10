package main

import (
	"errors"
	"fmt"
)

func main() {
	if message, err := VerifyName("Rocky Balboa"); err != nil { // HL
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}
}

// start error_handling OMIT
func VerifyName(name string) (message string, err error) { // HL
	if name != "Chiquinho Scarpa" {
		return message, errors.New("Erro: Seu nome não está na lista VIP")
	}

  message = fmt.Sprintf("Bem vindo Sr. %s. Gostaria de um vinho?", name)

  return // HL
}
// end error_handling OMIT
