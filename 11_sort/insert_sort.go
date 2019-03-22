package main

import "fmt"

//插入排序
func insertSort(a []int) []int{
	n := len(a)
	for i := 1; i < n ; i++  {
		val := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > val {
				a[j + 1] = a[j] //移动数据
			} else {
				break
			}

		}
		a[j + 1] = val
	}
	return a
}

func main()  {
	a :=  []int{8 , 5, 6, 1,3,2}
	newArr := insertSort(a)
	fmt.Println(newArr)
}
