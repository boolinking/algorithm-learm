package main

import "fmt"
//归并排序
func mergeSort(a []int) []int {
	n := len(a)
	if n < 2 {
		return a
	}
	q := n / 2
	left := mergeSort(a[:q])
	right := mergeSort(a[q:])
	return merge(left,right)
}

func merge(l []int , r []int) []int {
	tmp := make([]int , 0)
	i , j := 0 , 0
	for ; i < len(l) && j < len(r); {
		if l[i] > r[j] {
			tmp = append(tmp,r[j])
			j++
		} else {
			tmp = append(tmp,l[i])
			i++
		}
	}
	tmp = append(tmp,r[j:]...)
	tmp = append(tmp,l[i:]...)

	return tmp
}

func main()  {
	a := []int{1,8,5,3,7,2,4,6}
	a = mergeSort(a)
	fmt.Println(a)
}
