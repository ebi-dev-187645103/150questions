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

// s1 = xy = waterbottle
// x = wat
// y = terbottle
// s2 = yx = erbottlewat
// →s2はs1s1の部分一致に常になる。
func isRotation(s1,s2 string)bool{
	s1Len := len(s1)
	if s1Len==len(s2) && s1Len>0{
		s1s1 := s1+s1
		return strings.Contains(s1s1,s2)
	}
	return false
}

func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"

	fmt.Println(isRotation(s1,s2))
}
