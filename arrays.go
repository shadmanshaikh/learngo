package main

import (
	"fmt"
	"time"
)

func main(){

	start := time.Now()

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

	// lets perfom some crud operation on this array
	// creating new array
	oddNums := []int{}

	// appeneding elements 1 , 2
	oddNums = append(oddNums, 1 , 2 , 90 , 80)

	// reading array
	for i := 0  ; i < len(oddNums) ; i++ {
			fmt.Println(oddNums[i])
	}
	
	// again updating the array i.e oddNums =  [1 ,2]
	oddNums[1] = 3 // updated the element at position one
	fmt.Println(oddNums[1])
	
	// in range styled for loop 
	for _ , ele := range oddNums {
		fmt.Println(ele) 
	}

	// removing elements from the array using slicing
	oddNums = oddNums[1:]
	for _ , ele := range oddNums {
		fmt.Println(ele) 
	}

	end := time.Since(start)
	fmt.Print("Time Took to execute " , end)
	

}