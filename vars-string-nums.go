package main
import "fmt"

//global declaration 
var xan string = "hello , xan"

func main(){
	// ways of declaring & initializing a string
	var x string = "chadmax"
	var y = "gigachad"
	var a string
	a = "just declared"
	z := "gomax"
	fmt.Println(x , y  , z , xan , a) 

	// ways of using int , int8 , int16 , int32 , int64 , uint , uint8 ... 
	
	var num1 int = 8
	var num2 = 3 
	num3 := 4
	var num4 int8 = 122 // its in the range btw -128 to 127
	var num5 int16 = 442 // well you get the point higher the bit rate higher the number range
	
	fmt.Println(num1 , num2 , num3 , num4 , num5)

	// ways of using float32 , float64 
	var num6 float32 = 4.555
	var num7 float64 = 33422.41341198
	fmt.Println(num6 , num7)
}