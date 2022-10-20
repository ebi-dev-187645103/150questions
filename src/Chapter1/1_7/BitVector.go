// package main

// import (
// 	"fmt"
// 	"strconv"
// )
// func InsertStrToInt(str string)(val int){
// 	val,_ = strconv.Atoi(str)
// 	return val
// }
// 	func ConvertArrayStrToInt(str_ary []string)(int_ary []int){
// 	int_ary = make([]int, 0, len(str_ary))
// 	for _,str := range str_ary{
// 		int_ary = append(int_ary, InsertStrToInt(str))
// 	}
// 	return
// }

// func check(row []int)[]int{
// 	// bitVec := make([]uint64,(len(row)+63)/64)
// 	bitVec := 0
// 	fmt.Println(bitVec)
// 	for idx,val := range row{
// 		if val==0{
// 			bitVec = bitVec << idx
// 		}
// 	}
// 	return []int{0}
// }

// func main() {
// 	row := []int{1,0,2,3,4,5}
// 	serchZero := check(row)
// 	fmt.Println(serchZero)
// }
