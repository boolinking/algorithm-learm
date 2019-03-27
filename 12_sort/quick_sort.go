package main

import (
	"fmt"
)
//快排
func quickSort(a []int , start int , end int)  {
	if start >= end  {
		return
	}
	mid := partition(a,start,end)
	quickSort(a, start , mid-1)
	quickSort(a, mid+1,end)
}

func partition(a[] int, start int , end int) int {
	pivot := a[end]
	j := start
	for i := start; i < end ;i++  {
		if a[i] < pivot {
			a[i] ,a[j] = a[j] ,a[i]
			j++
		}
	}
	a[j] ,a[end] = a[end] ,a[j]
	return j
}

func main()  {
	a := []int{1,8,5,3,7,2,4,6}
	quickSort(a,0,len(a) -1)
	fmt.Println(a)
}
