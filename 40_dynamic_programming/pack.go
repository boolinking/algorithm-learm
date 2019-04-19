package main



func knapsack(weigth []int , n int , w int) int {
	var stat [n][w+1] bool
	stat[0][0] = true
	stat[0][weigth[0]] = true
	for i := 1; i < n ;i++  {
		for j := 0; j <= w ;j++  {
			if	stat[i-1][j] { //不放入i物品
				stat[i][j] = stat[i-1][j]
			}
		}

		for j := 0; j < w-weigth[i]; j++  {
			if stat[i-1][j] { //i放入背包
				stat[i][j+weigth[i]] = true
			}
		}
	}
	for i := w; i >= 0 ; i--  {
		if stat[n-1][i] {
			return i
		}
	}
	return 0
}

func knapsack2(weigth []int , n int , w int) int {
	var stat [w+1] bool
	stat[0] = true
	stat[weigth[0]] = true

	for i := 1; i < n ;i++  {
		for j := w - weigth[i]; j >= 0; j--  {
			if stat[j] {
				stat[j+weigth[i]] = true
			}
		}
	}

	for i := w; i >= 0 ; i--  {
		if stat[i] {
			return i
		}
	}
	return 0
}

func knapsack3(weight []int , value []int, n , w int) int {
	var stat [n][w+1]int
	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++  {
			stat[i][j] = -1
		}
	}
	stat[0][0] = 0
	stat[0][weight[0]] = value[0]

	for i := 1; i < n ;i++  {
		for j := 0; j <= w; j-- {
			if stat[i-1][j] >= 0 {//不放入
				stat[i][j] = stat[i-1][j]
			}
		}

		for j := w - weight[i]; j >= 0; j-- {
			if stat[i-1][j] >= 0 {//不放入
				v := stat[i-1][j] + value[i]
				if v > stat[i][j+weight[i]] {
					stat[i][j+weight[i]] = v
				}
			}
		}
	}
	maxVal := 0
	for i := 0; i <= w ; i++  {
		if stat[n-1][i] > maxVal {
			maxVal = stat[n-1][i]
		}
	}
	return maxVal

}



