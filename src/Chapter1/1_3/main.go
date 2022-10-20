package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)
func InsertStrToInt(str string)(val int){
	val,_ = strconv.Atoi(str)
	return val
}

func stringSort(str string) string{
	strAry := strings.Split(str, "")
	sort.Strings(strAry)
	return strings.Join(strAry,"")
}
func permutation(s,t string)bool{
	if len(s) != len(t){
		return false
	}
	sSort := stringSort(s)
	tSort := stringSort(t)
	return sSort == tSort
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
	s1 := "abcdefg"
	s2 := "dfagecb"
	s3 := "dfaged"
	s4 := "sdfajif"
	start := time.Now()
	fmt.Println("s1,s2:",permutation(s1,s2))
	end := time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
	fmt.Println("s1,s3:",permutation(s1,s3))
	fmt.Println("s1,s4:",permutation(s1,s4))
	fmt.Println("------------------")
	start = time.Now()
	fmt.Println("s1,s2:",permutation2(s1,s2))
	end = time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
	fmt.Println("s1,s3:",permutation2(s1,s3))
	fmt.Println("s1,s4:",permutation2(s1,s4))


}
