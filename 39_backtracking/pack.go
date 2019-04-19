package main

import "math"

//回溯实现0/1背包问题

var maxN int = math.MinInt64

var weight = []int{2,2,4,6,3} //物品的重量
var n int = 5 //物品的数量
var w int  = 9 //背包能承受的最大重量

func pkg(i,cw int)  {
	if i == 5 || cw == 9 { //物品装完了，或背包满了
		if maxN < cw {
			maxN = cw
		}
		return
	}

	pkg(i+1,cw) //不放入背包

	if cw + weight[i] < w {
		pkg(i+1,cw + weight[i]) //放入背包
	}
}

var stat [5][10]bool
func pkg2(i,cw int)  {
	if i == 5 || cw == 9 { //物品装完了，或背包满了
		if maxN < cw {
			maxN = cw
		}
		return
	}
	if stat[i][cw] {
		return
	}

	stat[i][cw] = true
	pkg(i+1,cw) //不放入背包

	if cw + weight[i] < w {
		pkg(i+1,cw + weight[i]) //放入背包
	}
}
