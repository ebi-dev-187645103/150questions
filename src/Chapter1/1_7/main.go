package main

import (
	"fmt"
	"strconv"
	"strings"
)
func InsertStrToInt(str string)(val int){
	val,_ = strconv.Atoi(str)
	return val
}
	func ConvertArrayStrToInt(str_ary []string)(int_ary []int){
	int_ary = make([]int, 0, len(str_ary))
	for _,str := range str_ary{
		int_ary = append(int_ary, InsertStrToInt(str))
	}
	return
}

func setZeros(matrix [][]int){
	row := make([]bool,len(matrix))
	col := make([]bool,len(matrix[0]))

	// 0を含む行及び列のインデックスを保存
	for y:=0;y<len(row);y++{
		for x:=0;x<len(col);x++{
			if matrix[y][x] == 0{
				row[y] = true
				col[x] = true
			}
		}
	}

	// 行iまたは行jが0を含んでいればarr[i][j]を0にする
	for y:=0;y<len(row);y++{
		for x:=0;x<len(col);x++{
			fmt.Println("------------------")
			fmt.Println("y:",y,"x:",x)
			fmt.Println("row[y]:",y,"row[x]:",x,row[y] || col[x])
			if row[y] || col[x]{
				matrix[y][x] = 0
			}
		}
	}
}

// func setZerosBitVector(matrix [][]int){
// 	row := make([]bool,len(matrix))
// 	col := make([]bool,len(matrix[0]))

// 	// 0を含む行及び列のインデックスを保存
// 	for y:=0;y<len(row);y++{
// 		for x:=0;x<len(col);x++{
// 			if matrix[y][x] == 0{
// 				row[y] = true
// 				col[x] = true
// 			}
// 		}
// 	}

// 	// 行iまたは行jが0を含んでいればarr[i][j]を0にする
// 	for y:=0;y<len(row);y++{
// 		for x:=0;x<len(col);x++{
// 			fmt.Println("------------------")
// 			fmt.Println("y:",y,"x:",x)
// 			fmt.Println("row[y]:",y,"row[x]:",x,row[y] || col[x])
// 			if row[y] || col[x]{
// 				matrix[y][x] = 0
// 			}
// 		}
// 	}
// }

func PrintField(field [][]int){
	fmt.Println("--------------------")
	for _,row := range field{
		fmt.Println(row)
	}
}

func Input()(field [][]int){
	input := []string{"0 1 2 3","4 5 6 7","8 9 10 11"}
	field = make([][]int,len(input))
	for idx,str := range input{
		strAry := strings.Split(str," ")
		strInt := ConvertArrayStrToInt(strAry)
		field[idx] = strInt
	}
	return
}
func main() {
	field := Input()
	PrintField(field)

	setZeros(field)
	PrintField(field)
}
