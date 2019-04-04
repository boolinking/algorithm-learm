package main

import "fmt"

//有序无重复数组，查找给定值
func binarySearch(a []int , k int) int {
	low := 0
	high := len(a) - 1
	for  low <= high  {
		mid := (high + low) / 2
		if a[mid] < k {
			low = mid + 1
		} else if a[mid] > k{
			high = mid - 1
		} else {
			return mid
		}

	}
	return -1
}


//查找第一个值等于给定值的元素
func binary(a []int , k int) int {
	low := 0
	high := len(a) -1
	for  low <= high  {
		mid := low + ((high - low) >> 1)
		if a[mid] < k {
			low = mid + 1
		} else if a[mid] > k{
			high = mid - 1
		} else {
			if mid == 0 || a[mid-1] != k {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1

}

//查找最后一个值等于给定值的元素
func binary2(a []int , k int) int {
	low := 0
	high := len(a) -1
	for  low <= high  {
		mid := low + ((high - low) >> 1)
		if a[mid] < k {
			low = mid + 1
		} else if a[mid] > k{
			high = mid - 1
		} else {
			if mid == len(a) -1 || a[mid+1] != k {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

// 查找第一个大于等于给定值的元素
func binary3(a []int , k int) int {
	low := 0
	high := len(a) -1
	for  low <= high  {
		mid := low + ((high - low) >> 1)
		if a[mid] < k {
			low = mid + 1
		} else if a[mid] >= k{
			if mid == 0 || a[mid-1] < k {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

//查找最后一个小于等于给定值的元素
func binary4(a []int , k int) int {
	low := 0
	high := len(a) -1
	for  low <= high  {
		mid := low + ((high - low) >> 1)
		if a[mid] > k {
			high = mid - 1
		} else if a[mid] <= k{
			if mid == len(a) -1 || a[mid+1] > k {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}






func main()  {
	a := []int{1,3,4,5,6,8,8,8,10,18}
	n := binary4(a,8)
	fmt.Println(n)
}