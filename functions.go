package main

import "fmt"

func addNum(a int, b int) int {
	sum := a + b
	return sum
}

// function to find the smalles element in an array 

func smallEle(arr []int) int {
    if len(arr) == 0 {
        panic("Array cannot be empty") // Handle empty arrays
    }

    smallestEle := arr[0] // Initialize with the first element
    for i := 1; i < len(arr); i++ {
        if arr[i] < smallestEle {
            smallestEle = arr[i]
        }
    }
    return smallestEle
}

//

func main() {
	// taking input from user
	var a , b int
	println("Enter the numbers to add : ")
	fmt.Scan(&a , &b)
	fmt.Println("The sum is :" , addNum(a , b))

	//finding the smallest ele
	newArr := []int{1 , 2 , 3}

	fmt.Println("The smalles element is : " , smallEle(newArr))

}