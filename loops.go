package main

import (
	"fmt" 
	"sort"
	)

func main(){
	// playing around with loops
	arr := []int{1, 89,391,90 ,12}
	for i:=0 ; i < len(arr) ; i++ {
		fmt.Println(arr[i])
	}
	println("Sorting the array using sort()")
	sortedArr := []int{}
	sort.Ints(arr)
	sortedArr = append(sortedArr, arr...)
	for i := 0 ; i < len(sortedArr) ; i++ {
		println(sortedArr[i])
	}

	//finding the largest ele in the array

	chadArr := []int{1,2,8,12,4,34,234}

	largestEle := 0

	for i := 0 ;i < len(chadArr) ; i++{
		if(largestEle < chadArr[i]){
			largestEle = chadArr[i]
		}
	}
	println("The largest element is :")
	println(largestEle)

	//binary search in go lang

	

	// append(sortedArr, [])
}