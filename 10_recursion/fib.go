package main

import "fmt"

func Fib(n int)  int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n - 2)
}
//递归实现
func multi(n int) int {
	if n == 1 {
		return 1
	}
	return n * multi(n - 1)
}

//尾递归实现
func mult(n int, result int) int {
	if n == 1 {
		return result
	}
	return mult(n - 1 , result*n)
}

func main()  {
	//fmt.Print(Fib(8))
	fmt.Println(multi(5))
	fmt.Println(mult(5,1))
}