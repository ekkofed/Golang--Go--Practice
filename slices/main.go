package main
import (
	"fmt"
	"math/rand"
)

func main()  {
	//SLICES
	fmt.Println("Define slices")
	var numbers[] int
	numbers = make([]int,5)
	matrix:= make([][]int,3*3)

	//Insert data

	fmt.Println(">>>>Insert slice data")
	for i:=0;i<5;i++{
		numbers[i]= rand.Intn(100) //random number

	}

	//Insert matrix data
	fmt.Println(">>>>>> Insert slice matrix data")
	for i:=0;i<3;i++{
		matrix[i]= make([]int,3)
		for j:=0;j<3;j++{
			matrix[i][j] = rand.Intn(100)
		}
	}

	//display the data
	fmt.Println(">>>>> Displaying slice data")
	for j:=0;j<5;j++{
		fmt.Println(numbers[j])
	}

	//Display matrix data
	fmt.Println(">>>>>> Displaying matrix slice data 2D")
	for i:=0;i<3;i++{
		for j:=0; j<3;j++{
			fmt.Println(matrix[i][j])
		}
	}
}