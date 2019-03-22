package main

import "fmt"
//选择排序
func selectSort(a []int) []int  {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i; j < n - 1 ; j++  {
			if a[i] > a[j + 1] {
				a[i] , a[j+1] = a[j + 1] , a[i]
			}
		}
	}
	return a
}

func main()  {
	a :=  []int{8 , 5, 6, 1,3,2,7}
	newArr := selectSort(a)
	fmt.Println(newArr)
}