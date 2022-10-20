package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)
func InsertStrToInt(str string)(val int){
	val,_ = strconv.Atoi(str)
	return val
}

func replaceSpaces(str string) string{
	var b bytes.Buffer
	for idx, r := range str {
		fmt.Println(str[idx:idx+1],r)
		if r == 32 { // " " == 32
			b.WriteString("%20")
		}else{
			b.WriteRune(r)
			}
	}
	return b.String()

}

func permutation2(s,t string)bool{
	if len(s) != len(t){
		return false
	}

		letters := make([]int,256)
		for i:=0;i<len(s);i++{
			tmp1 := []byte(s[i:i+1])
			tmp1Int,_ := strconv.Atoi(fmt.Sprintf("%d",tmp1[0]))
			letters[tmp1Int]++
		}

		for i:=0;i<len(t);i++{
			tmp1 := []byte(t[i:i+1])
			tmp1Int,_ := strconv.Atoi(fmt.Sprintf("%d",tmp1[0]))
			letters[tmp1Int]--
			if letters[tmp1Int] < 0{
				return false
			}
		}
		return true
	}


func main() {
	s1 := "a b c d "
	start := time.Now()
	fmt.Println("s1:",replaceSpaces(s1))
	end := time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
	fmt.Println("====================")
	start = time.Now()
	fmt.Println("s1:",strings.ReplaceAll(s1," ","%20"))
	end = time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
}
