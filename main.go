package main

import(
	"fmt"
)

var i = 45

func main()  {

	var i = 55
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

	fmt.Printf("value i  : %v\n", i)

	for true{
		i:=11
		fmt.Printf("value : %v\n", i)
		i=12
		fmt.Printf("value : %v\n", i)
		break
	}

	toto()
} 

func toto (){
	fmt.Printf("value i  : %v\n", i)
}