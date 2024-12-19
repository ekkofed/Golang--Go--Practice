package main

import "fmt"

func main()  {
	var i int

	for i=0;i<5;i++{
		fmt.Println(i)
	}

	for j:=5;j<11;j++{
		fmt.Println(j)
	}


	// Iteration while
	/*Go doesn't provide while syntax. we can use for */
	f:=0
	for f<5{
		fmt.Println(f)
		f++
	}

	//BREAK AND CONTINUE

	var k int
	for k=0;k<5;k++{
		if k==3{
			break
		}
		fmt.Println(k)
	}
	for w:=5;w<11;w++{
		if w==7{
			continue
		}
		fmt.Println(w)
	}
}