package main

import "fmt"

//求无序数组第K大元素 时间复杂度为O(n)
func search(a []int,k int, start int , end int)  int {
	if start >= end {
		return a[start]
	}
	p := partitions(a,start,end)
	i := p + 1
	if  k < i {
		return search(a , k , start , p - 1)
	} else if k >  i {
		return search(a , k , i , end)
	} else {
		return a[p]
	}


}

func partitions(a []int , start int , end int) int {
	p := a[end]
	i := start
	for j := start; j < end ;j++  {
		if a[j] > p {
			a[i] , a[j] = a[j] , a[i]
			i++
		}
	}
	a[i] , a[end] = a[end] , a[i]
	return i
}



func main()  {
	a := []int{1,8,5,3,7,2,4,6}
	v := search(a,3,0,len(a) -1)
	fmt.Println(v)
}
