package main

import(
	"fmt"
)

var i = 45
var a =100

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
	pointer()

	fmt.Printf("scop a 1 : %v\n", a)
	  a := 200
	fmt.Printf("scop a 2 : %v \n", a)
} 

func toto (){
	fmt.Printf("value i  : %v\n", i)
}



func pointer(){
	//This is because receipt and receipt2 pointing to the same coat.
//The operator & means “give me the address of the variable”. So this is de-referecing. The memory address ist the place where the data is stored.
//The operator * means tread this as a pointer. So *receipt fetches the address stored in receipt and uses the content which is stored there.
	coat := "I am a coat"
	receipt := &coat
	receipt2 := &coat  
	fmt.Println(*receipt)      
	fmt.Println(*receipt2)
  
}