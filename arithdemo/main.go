package main
import "fmt"

func main(){
	var a,b int
	a=5
	b=10

	//arithmetic operations

	//addition
	c:=a+b 
	fmt.Printf("%d + %d = %d \n",a,b,c)

	//subtraction
	d:=a-b
	fmt.Printf("%d - %d = %d\n",a,b,d)

	//division
	e:= float32(a) / float32(b)
	fmt.Printf("%d / %d = %.2f \n",a,b,e) 

	f:=a*b 
	fmt.Printf("%d * %d = %d \n",a,b,f)

}