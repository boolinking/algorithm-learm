package main

import "math"

var minDist int = math.MaxUint64
func minDistBT(i , j , dist int , w [][]int , n int)  {
	if i == n && j == n {
		if dist < minDist {
			minDist = dist
		}
		return
	}
	if j < n {//往左走
		minDistBT(i,j+1,dist+w[i][j],w,n)
	}
	if i < n {//往下走
		minDistBT(i+1,j,dist+w[i][j],w,n)
	}
}
