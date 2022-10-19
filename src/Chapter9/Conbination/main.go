package main

import (
	"fmt"
	"strconv"
	"time"
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

func Input()(n,r int){
	n = 4
	r = 2

	return
}

func conbination(n,r int)int{
	nCr := 1
	for i:=1;i<=r;i++{
		nCr = nCr*(n-i+1)/i
	}
	return nCr
}

func conbination1(n,r int)int{
	if r==0 || r==n{
		return 1
	}else{
		return conbination1(n-1,r-1)+conbination1(n-1,r)
	}}

func main() {
	n,r := Input()

	start := time.Now()
	nCr := conbination(n,r)
	end := time.Now()
	fmt.Println(nCr)
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
	fmt.Println("--------------------")

	start1 := time.Now()
	nCr1 := conbination1(n,r)
	fmt.Println(nCr1)
	end1 := time.Now()
	fmt.Printf("%f秒\n",(end1.Sub(start1)).Seconds())
}
