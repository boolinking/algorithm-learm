package main

import "fmt"

//冒泡排序
func bubbleSort(a []int) []int  {
	n := len(a)
	fmt.Println(n)
	for i := 0; i < n ; i++  {
		var flag bool = false
		for j := 0; j < n- i - 1 ; j++  {
			if a[j] > a[j + 1] {
				a[j] , a[j + 1] = a[j + 1] , a[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return a
}

func main()  {
	a :=  []int{8 , 5, 6, 1,3,2}
	newArr := bubbleSort(a)
	fmt.Println(newArr)
}
