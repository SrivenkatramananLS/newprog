package main

import "fmt"

func greet() {
	nums0 := [3]int{1,3,4} //array
	nums1 := []int{7,8,8}  // slices
    nums1 = append(nums1, 40)
	fmt.Println(nums0,nums1)
}

func mathematical(){
	numb := []int{1,2,3,4,5}

	for i := range numb{
		numb[i] = numb[i] * 2
		//fmt.Println(numb)
	}
	fmt.Println(numb)
}

func tmp(){
	temp := [7]int{31,30,40,43,41,33}
	fmt.Println(temp)

}


