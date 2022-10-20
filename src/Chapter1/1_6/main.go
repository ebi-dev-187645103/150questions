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
func rotate(matrix [][]string){
	side := len(matrix)
	for layer:=0;layer<side/2;layer++{
		first := layer
		last  := side - 1 - layer
		for i:=first;i<last;i++{
			offset:= i-first
			fmt.Println("----")
			fmt.Println("first :",first)
			fmt.Println("last  :",last)
			fmt.Println("i     :",i)
			fmt.Println("offset:",offset)
			fmt.Println("上端:",first,i)
			fmt.Println("左端:",last-offset,first)
			fmt.Println("下端:",last,last-offset)
			fmt.Println("右端:",i,last)

			// 上端を保存
			top:= matrix[first][i]
			// 上端←左端
			matrix[first][i] = matrix[last-offset][first]
			// 左端←下端
			matrix[last-offset][first] = matrix[last][last-offset]
			// 下端←右端
			matrix[last][last-offset] = matrix[i][last]
			// 右端←上端
			matrix[i][last] = top
		}
	}
}
func PrintField(field [][]string){
	fmt.Println("--------------------")
	for _,row := range field{
		fmt.Println(row)
	}
}

func Input()(field [][]string){
	input := []string{"a b c d","e f g h","i j k l","m n o p"}
	field = make([][]string,len(input))
	for idx,str := range input{
		strAry := strings.Split(str," ")
		field[idx] = strAry
	}
	return
}
func main() {
	field := Input()
	PrintField(field)

	rotate(field)
	PrintField(field)
}
