package main
import "fmt"

func main(){
	// ways of using int arrays 
	 x  := [3]int{1,2,3}
	fmt.Println(x)
	x[1] = 15 //changed the element at position 1 i.e 2 -> 15
	fmt.Println(x)

	//ways of using string arrays

	big4lCommunity := [3]string{"chadmax" , "gigachad" , "xan"}
	fmt.Println(big4lCommunity)

	//ways of using slice [elements can be appended to it]

	nums := []int{3 , 67 , 12 , 93 , 45 }
	nums = append(nums, 99 , 78 ,23 ,93)
	fmt.Println(nums)
	
}