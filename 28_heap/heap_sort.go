package main

import "fmt"

func heapSort(a [] int , n int)  {

	buildHeap(a , n)
	k := n
	for k > 1 {
		a[1] , a[k] = a[k] , a[1]
		Heapify(a, 1, k-1)
		k--
	}

}

func buildHeap( a []int , n int )  {
	for i := n / 2; i >= 1; i-- {
		Heapify(a, i, n)
	}
}


func Heapify(a []int, i int , n int)  {
	for {
		maxPos := i
		if i*2 <= n && a[i] < a[2*i] {
			maxPos = 2 * i
		}
		if (2*i + 1) <= n && a[maxPos] < a[2 * i + 1]  {
			maxPos = 2 * i +1
		}
		if maxPos == i {
			break
		}
		a[i] , a[maxPos] = a[maxPos] , a[i]
		i = maxPos
	}
}

func main()  {
	a := []int{1,4,2,9,5,6,8,7}
	heapSort(a,len(a) - 1)
	fmt.Println(a)
}