package main

import "fmt"

type Heap struct {
	a []int
	capacity int
	count int
}

func New(capacity int) *Heap {
	return &Heap{make([]int,capacity),capacity,0}
}

func (heap *Heap) Insert(data int)  {//大顶堆

	if heap.count == heap.capacity {//堆满了
		return
	}
	heap.count++
	heap.a[heap.count] = data
	i := heap.count
	for i / 2 > 0 && heap.a[i] > heap.a[i / 2] {//从下往上堆化
		//如果比父节点大就交换位置
		heap.a[i] , heap.a[ i / 2] = heap.a[ i / 2],heap.a[i]
		i = i / 2
	}

}

//删除堆最大元素(数组位置1的元素)
func (heap *Heap) Del()  int {

	if heap.count == 0 {
		return -1
	}
	data := heap.a[1]
	//为保持数组连续（完全二叉树形态）将末尾元素放到1的位置然后堆化
	heap.a[1] = heap.a[heap.count]
	heap.count--
	heapify(heap.a,1,heap.count)
	return data
}

func (heap *Heap) Sort() []int {
	k := heap.count
	for k > 1  {
		heap.a[1] , heap.a[k] = heap.a[k] , heap.a[1]
		heapify(heap.a,1, k)
		k--
	}

	return heap.a
}
//从上往下堆化
func heapify(a []int, i int , n int)  {
	for {
		maxPos := i
		if 2*i <= n && a[i] < a[2*i] {
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
	heap := New(10)
	heap.Insert(5)
	heap.Insert(8)
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(1)

	fmt.Println(heap.Del())
	fmt.Println(heap.Del())
	fmt.Println(heap.Del())
	fmt.Println(heap.Del())
	fmt.Println(heap.Del())

	heap.Insert(5)
	heap.Insert(8)
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(1)
	fmt.Println(heap.Sort())
}

