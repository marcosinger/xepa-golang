package main

import "fmt"

// start struct OMIT
type Fighter struct { // HLStruct
	Name string
	Age  int
}

// end struct OMIT

// start method OMIT
func (f Fighter) Greeting() { // HL
	fmt.Printf("Ele vem ai! %s/%d", f.Name, f.Age)
}

// end method OMIT

func main() {
	// start struct apollo OMIT
	apollo := Fighter{"Apollo Creed", 35} // importa a ordem e a quantidade de parâmetros
	// end struct apollo OMIT

	// start struct rocky OMIT
	rocky := Fighter{Age: 25, Name: "Rocky Balboa"} // a ordem não importa
	// end struct rocky OMIT

	// start struct yoda OMIT
	yoda := Fighter{Name: "Mestre Yoda"} // parâmetro Age omitido
	// end struct yoda OMIT

	// start print struct OMIT
	fmt.Printf("Fighter Apollo: %v\n", apollo)
	fmt.Printf("Fighter Rocky: %#v\n", rocky)
	fmt.Printf("Fighter Yoda: Nome -> %s | Idade -> %d\n", yoda.Name, yoda.Age)
	// end print struct OMIT

	// start print method OMIT
	rocky.Greeting()
	// end print method OMIT
}
