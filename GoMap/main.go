package main

import (
	"fmt"
	"math/rand"
)
func main()  {
	
	//define array
	fmt.Println("Define Map")
	products :=make(map[string]int)
	customers:= make(map[string]int)

	//Insert data
	fmt.Println(">>>>>>> insert map data")
	products["Product1"] = rand.Intn(100)
	products["Product2"] = rand.Intn(100)

	customers["cust1"]= rand.Intn(100)
	customers["cust2"]= rand.Intn(100)

	//Display data
	fmt.Println(">>>>>>>> display map data")
	fmt.Println(products["product1"])
	fmt.Println(products["product2"])
	fmt.Println(customers["cust1"])
	fmt.Println(customers["cust2"])
}