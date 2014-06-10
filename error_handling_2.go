package main

import (
	"errors"
	"fmt"
)

// start error_handling OMIT
func main() {
	if message, err := VerifyName("Rocky Balboa"); err != nil { // HL
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}
}

// end error_handling OMIT

func VerifyName(name string) (string, error) {
	if name != "Chiquinho Scarpa" {
		return "", errors.New("Erro: Seu nome não está na lista VIP")
	}

	return fmt.Sprintf("Bem vindo Sr. %s. Gostaria de um vinho?", name), nil
}
