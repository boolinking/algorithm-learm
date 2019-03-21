package _5_array

import (
	"errors"
	"fmt"
)



type Array struct {
	data []int
	len uint
	cap uint
}

func New(capacity uint) *Array {
	return &Array{
		data:make([]int,capacity),
		len:0,
		cap:capacity,
	}
}

func (arr *Array) Len() uint  {
	return arr.len
}

func (arr *Array) Cap() uint {
	return arr.cap
}



//数组越界检查
func (arr *Array) IndexOutOfArray(index uint) bool {
	if index >= arr.Cap() {
		return true
	}
	return false
}


//在指定位置插入数据
func (arr *Array) Insert(index uint, data int) error {
	if ok := arr.IndexOutOfArray(index); ok {
		return errors.New("out of index 05_array")
	}
	if arr.Len() == arr.Cap() {
		return errors.New("05_array full")
	}
	for i := arr.Len(); i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[index] = data
	arr.len++
	return nil
}

//在指定位置删除一个元素
func (arr *Array) Delete(index uint) (int , error) {
	if ok := arr.IndexOutOfArray(index); ok {
		return 0,errors.New("index of 05_array")
	}
	v := arr.data[index]
	for i:= index; i < arr.Len()-1; i++  {
		arr.data[i] = arr.data[i + 1]
	}
	arr.len--
	return v,nil
}

//随机访问指定索引的元素
func (arr *Array) Find(index uint) int {
	if ok:= arr.IndexOutOfArray(index); ok {
		return -1
	}
	return arr.data[index]
}

//打印数组
func (arr *Array) Print()  {
	for i:= range arr.data {
		fmt.Printf(" %d",arr.data[i])
	}
	fmt.Println()
}