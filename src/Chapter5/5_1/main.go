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

func updateBits(n,m,i,j int)int{
	// すべてが1のビット列
	allOnes := ^0

	//jより前が1で後が0の列。left=11100000
	left := allOnes << (j+1)

	// iより後が1の列。eight = 00000011
	right := ((1 << i)-1)

	// iとjの間は0でそれ以外は1の列。mask = 11100011
	mask := left | right

	// jビット目からiビット目をクリアし、mを挿入する
	n_cleared  := n & mask
	m_shifteed := m << i

	return n_cleared | m_shifteed
}

func main() {
	// 入力:n=10000000000(2),m=1011(2),i=2,j=6
	// 出力:n=10001001100(2)
	fmt.Println(updateBits(1024,19,2,6))
}
