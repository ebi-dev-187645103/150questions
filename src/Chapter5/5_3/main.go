package main

import (
	"bytes"
	"fmt"
	"strconv"
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

func PrintBinary(num float64)string{
	if num>=1 || num<=0{
		return "ERROR"
	}
	var b bytes.Buffer
	b.WriteString(".")
	for num>0 {
		if len(b.String()) >= 32 {
			fmt.Println(b.String())
			return "ERROR"
			// return b.String()
		}

		r:=num*2
		if r>= 1{
			b.WriteString("1")
			num = r-1
		}else{
			b.WriteString("0")
			num = r
		}
	}
	return b.String()
}

func PrintBinary2(num float64)string{
	if num>=1 || num<=0{
		return "ERROR"
	}
	var b bytes.Buffer
	frac := 0.5
	b.WriteString(".")
	for num>0{
		if len(b.String()) >= 32 {
			fmt.Println(b.String())
			return "ERROR"
			// return b.String()
		}
		if num >= frac{
			b.WriteString("1")
			num -= frac
		}else{
			b.WriteString("0")
		}
		frac/=2
	}
	return b.String()
}


func main() {
	num := 0.1
	fmt.Println(PrintBinary(num))
	fmt.Println(PrintBinary2(num))
}
