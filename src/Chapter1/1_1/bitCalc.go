package main

import "fmt"
func main() {
	a := 10 // 1010
	b := 12 // 1100

	// a&b
	// a = 1010 = 10
	// b = 1100 = 12
	//---------------
	//     1000 = 8
	bitAnd := a&b
	fmt.Println("a&b :",bitAnd)

	// a|b
	// a = 1010 = 10
	// b = 1100 = 12
	//------------------
	//     1110 = 14
	bitOr  := a|b
	fmt.Println("a|b :",bitOr)

	// a^b
	// a = 1010 = 10
	// b = 1100 = 12
	//------------------
	//     0110 = 6
	bitXor := a^b
	fmt.Println("a^b :",bitXor)

	d := 4             // 0100 = 4
	leftShift  := d<<1 // 1000 = 8
	fmt.Println("4<<1 :",leftShift)
	rightShift := d>>1 // 0010 = 2
	fmt.Println("4>>1 :",rightShift)

	//   10 = 1 を2bit左シフト
	// 1000 = 8
	c := 1 << 2
	fmt.Println("1<<2 :",c)
}
