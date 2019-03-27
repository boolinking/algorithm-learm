package main

import "fmt"

//计数排序
func countSort(a []int)  {
	n := len(a)

	max := a[0]

	//求数组最大值
	for i := 0; i < n ; i++  {
		if a[i] > max {
			max = a[i]
		}
	}
	//申请一个数组c
	c := make([]int,max+1)

	//统计数组a中每个元素出现的次数放于数组c
	for i := 0; i < n ; i++  {
		c[a[i]]++ //c数组的下标为a数组的元素，值为元素出现的次数
	}

	//累加
	for i := 1; i <= max ; i++  {
		c[i] = c[i-1] + c[i] //c[i]的值为a中小于等于 i的个数
	}


	r := make([]int,n)

	//
	for i := 0; i < n; i++  {
		index := c[a[i]] -1
		r[index] = a[i]
		c[a[i]]--
	}

	//将数据复制到a
	for i := 0; i < n ; i++  {
		a[i] = r[i]
	}

}



func main()  {
	a := []int{1,8,5,3,7,2,4,6}
	countSort(a)
	fmt.Println(a)
}
