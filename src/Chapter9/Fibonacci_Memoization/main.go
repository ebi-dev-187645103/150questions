package main

import "fmt"

var memo [11]int

//http://blog.kzfmix.com/entry/1315391392
func fib(n int) int {
	fmt.Println("-------------------")
	fmt.Println("n=",n)
	if n <= 1 {
		return n
	}
	fmt.Println("memo[n]=",memo[n])
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fib(n-1) + fib(n-2)
	fmt.Println("|------------------")
	fmt.Println("n =",n,"memo[n] =",memo[n])
	return memo[n]
}

func main() {
	fmt.Println(fib(10))
}
