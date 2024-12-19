package main

import "fmt"

func main()  {
	
	var (
		a=5
		b=8
	)

	if a>b || a-b<a{
		fmt.Println("Conditional --> a>b || a-b <a ")
	}else{
		fmt.Println("... another")
	}


	// SWITCH CASE

	fmt.Println("Enter the selected number: ")
	var selected int
	fmt.Scanf("%d",&selected)

	switch selected {
	case 0:
		fmt.Println("Selected = 0")

	case 1:
		fmt.Println("Selected = 1")

	case 2:
		fmt.Println("Selected = 2")

	case 3:
		fmt.Println("Selected = 3")
	default:
		fmt.Println("other....")
	}

}