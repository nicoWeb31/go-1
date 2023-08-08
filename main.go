package main

import(
	"fmt"
)

func main()  {
	varInt := 12
	varString :="toto"
	tata := "tata"
	varFloat := "3.14"

	var titi  = "titi"
	

	fmt.Printf("hello go")
	fmt.Printf("value : %v\n", varInt)
	fmt.Printf("type: %T\n", varInt)
	fmt.Printf("value : %v\n", varString)
	fmt.Printf("type : %T\n", varString)
	fmt.Printf("value : %v\n", tata)
	fmt.Printf("value tata : %v\n", titi)
	fmt.Printf("value : %v\n", varFloat)
	fmt.Printf("Type : %T\n", varFloat)

}    