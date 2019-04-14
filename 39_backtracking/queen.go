package main

import "fmt"

//下标为表示行，值表示 queen存储的列
var result [8]int

func cal8queen(row int)  {
	if row == 8 {
		print()
		return
	}
	for column := 0; column < 8 ; column++  {
		if isOk(row,column) {
			result[row] = column
			cal8queen(row+1)
		}
	}
}

func isOk(row , column int) bool {
	leftUp , rigntUp := column -1 , column +1
	for i := row - 1; i >= 0 ; i--  {
		if result[i] == column {
			return false
		}
		if leftUp >= 0 {
			if result[i] == leftUp {
				return false
			}
		}
		if rigntUp < 8 {
			if result[i] == rigntUp {
				return false
			}
		}
		leftUp--
		rigntUp++
	}
	return true
}

func print()  {
	i := 0
	for row := 0; row < 8 ; row++  {
		for col := 0 ; col < 8 ; col++  {
			if result[row] == col {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}

		}
		fmt.Println()
	}
	i++
	fmt.Println()
}

func main()  {
	cal8queen(0)

}