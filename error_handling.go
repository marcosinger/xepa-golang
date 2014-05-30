package main

// start error_handling OMIT
import (
	"errors" // HL
	"fmt"
)

func main() {
	message, err := VerifyName("Rocky Balboa") // HL

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}
}

func VerifyName(name string) (string, error) { // HL
	if name != "Eike Batista" {
		return "", errors.New("Erro: Seu nome não está na lista VIP")
	}

	return fmt.Sprintf("Bem vindo Sr. %s. Gostaria de um vinho?", name), nil
}

// end error_handling OMIT
