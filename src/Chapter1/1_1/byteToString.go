package main

import (
	"fmt"
)
func main() {
	tmp := "abcd"
	t := tmp[0:1]
	fmt.Println(t)
	idx := []byte(t)
	fmt.Println(idx)
	c :=idx[0]
	fmt.Println(c)
	fmt.Println("------------------")
	a := fmt.Sprintf("%s,%x,%d",idx,idx,idx)
	fmt.Println(a)
	b := fmt.Sprintf("%d",idx[0])
	fmt.Println(b)
}
