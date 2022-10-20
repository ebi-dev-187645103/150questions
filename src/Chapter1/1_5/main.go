package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)
func InsertStrToInt(str string)(val int){
	val,_ = strconv.Atoi(str)
	return val
}

func compress1(str string) string{
	// var buffer bytes.Buffer
	compStr := ""
	last := str[0:1]
	continueNum := 1
	for i:=1;i<len(str);i++{
		target := str[i:i+1]
		if target == last{
			continueNum++
		}else{
			if continueNum == 1{
				compStr += last
			}else{
				compStr += strconv.Itoa(continueNum) + last
			}
			last = target
			continueNum=1
		}
	}

	return compStr
}

func compress2(str string) string{
	var b bytes.Buffer
	var last rune
	continueNum := 0
	// for i:=0;i<len(str);i++{
	for idx,r := range str{
		if idx==0{
			last = r
		}
		if r == last{
			continueNum++
		}else{
			if continueNum == 1{
				b.WriteRune(r)
			}else{
				b.WriteString(strconv.Itoa(continueNum))
				b.WriteRune(r)
			}
			last = r
			continueNum=1
		}
	}

	return b.String()
}

func main() {
	s1 := "aabcccccaaa"
	fmt.Println(s1)
	fmt.Println("====================")
	start := time.Now()
	fmt.Println("s1:",compress1(s1))
	end := time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
	fmt.Println("====================")
	start = time.Now()
	fmt.Println("s1:",compress2(s1))
	end = time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	fmt.Println("----------------------------")

	s2 := "aabcccccaaannnnnnnniiiiiiiiiiiiiiiiiiiiiiiibbbbbbbbbbbbbbbbbbbbbbbbbpppppppppppppppssssssssjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjbbbbbbbbbbbbbbbbbbbbbpodfshoijfskvm;jgrissssssssssccccccccccccccccccccccc5555555555555555"
	fmt.Println(s1)
	fmt.Println("====================")
	start = time.Now()
	fmt.Println("s2:",compress1(s2))
	end = time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
	fmt.Println("====================")
	start = time.Now()
	fmt.Println("s2:",compress2(s2))
	end = time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
}
