package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println("Circle Area Calculation")
	fmt.Print("Enter a radius value: ")

	var radius float64
	fmt.Scanf("%f",&radius)

	area:=math.Pi * math.Pow(radius,2)
	fmt.Printf("Circle area with radius %.2f = %.2f \n",radius,area)


	// Comparison operators
	var(
		b=2
		a=10
	)
	fmt.Println(a>b)
	fmt.Println(a<b)
	fmt.Println(a>=b)
	fmt.Println(a<=b)
	fmt.Println(a!=b)
	fmt.Println(a==b)


	/*LOGICAL OPERATORS-----They are used for determining the logic between
	variables or values*/

	var(
		c=5
		d=8
	)
	fmt.Println(c>d && c!=d)
	fmt.Println(!(c>=d))
	fmt.Println(c==d || c>b)
}