package main

import "math"

func minDistBT(metrix [][]int , n int){
	status := make([][]int,n)
	sum := 0
	for i :=0; i < n; i++ {
		sum += metrix[i][0]
		status[i][0] = sum
	}
	sum = 0
	for j :=0; j < n; j++ {
		sum += metrix[0][j]
		status[0][j] = sum
	}

	for i := 1; i < n ;i++ {
		for j := 1; j < n; j++ {
			status[i][j] = metrix[i][j] + min(status[i-1][j],status[i][j-1])
		}
	}
}

func min(a , b int) int {
	if a > b {
		return a
	}
	return b
}

var metrix = [4][4]int {{1,3,5,9}, {2,1,3,4},{5,2,6,1},{8,6,4,3}}
var mem  [4][4]int
var n int = 4

func minDistBT2(i ,j int) int {
	if i == 0 && j == 0 {
		return metrix[0][0]
	}
	if mem[i][j] > 0 {
		return mem[i][j]
	}

	minUp := math.MaxUint64
	if i >= 0 {
		minUp = minDistBT2(i-1,j)
	}

	minLeft := math.MaxUint64
	if j >= 0 {
		minLeft = minDistBT2(i,j-1)
	}
	currMin := metrix[i][j] + min(minLeft,minUp)
	mem[i][j] = currMin
	return currMin
}


