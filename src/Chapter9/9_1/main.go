package main

import (
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

func Input()(stairs int){
	stairs = 4

	return
}

func countWays(n int)int{
	// fmt.Println("-------------------")
	// fmt.Println("n=",n)
	if n<0{
		return 0
	}
	if n == 0{
		return 1
	}
	return countWays(n-1)+countWays(n-2)+countWays(n-3)
}

var memo [11]int
// n=37でオーバーフローするため、それ以上の場合は、型を変更すること。
func countWaysDP(n int)int{
	fmt.Println("-------------------")
	fmt.Println("n=",n)
	fmt.Println("memo[n]=",memo[n])

	if n < 0 {
		fmt.Println("n < 0")
		return 0
	}
	if n == 0{
		fmt.Println("n == 0")
		return 1
	}
	// if memo[n] > -1 {
	// 	fmt.Println("memo[n] > -1")
	if memo[n] > 0 {
		fmt.Println("memo[n] > 0")
		return memo[n]
	}
	memo[n] = countWaysDP(n-1) + countWaysDP(n-2) + countWaysDP(n-3)
	fmt.Println("|------------------")
	fmt.Println("n =",n,"memo[n] =",memo[n])
	return memo[n]
}

func main() {
	stairs := Input()

	// start1 := time.Now()
	// patterns1 := countWays(stairs)
	// end1 := time.Now()
	// fmt.Printf("%f秒\n",(end1.Sub(start1)).Seconds())
	// fmt.Println(patterns1)
	// fmt.Println("--------------------")

	// start2 := time.Now()
	// fmt.Println(memo)
	patterns2 := countWaysDP(stairs)
	// end2 := time.Now()
	// fmt.Printf("%f秒\n",(end2.Sub(start2)).Seconds())
	fmt.Println(patterns2)

	// patterns := countWaysDP(stairs,)
}
