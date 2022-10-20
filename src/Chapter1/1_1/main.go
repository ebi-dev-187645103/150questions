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

func isUniqChars(str string)bool{
	// 文字の種類は256文字しかないため、それより多い文字数にユニークな文字列はあり得ない。
	if len(str) >256{
		return false
	}

	checker := 0
	for i:=0;i<len(str);i++{
		tmp1 := []byte(str[i:i+1])
		tmp1Int,_ := strconv.Atoi(fmt.Sprintf("%d",tmp1[0]))
		tmp2 := []byte("a")
		tmp2Int,_ := strconv.Atoi(fmt.Sprintf("%d",tmp2[0]))
		val := tmp1Int - tmp2Int

		if (checker & (1 << val)) > 0 {
			return false
		}
		checker |= (1 << val)
	}
	return true
}

func isUniqChars2(str string)bool{
	// 文字の種類は256文字しかないため、それより多い文字数にユニークな文字列はあり得ない。
	if len(str) >256{
		return false
	}

	charSet := make([]bool,256)
	for i:=0;i<len(str);i++{
		idxBytes := []byte(str[i:i+1])
		if len(idxBytes)!=1{
			fmt.Println("over")
			return false
		}
		idxStr := fmt.Sprintf("%d",idxBytes[0])
		idxInt,_ := strconv.Atoi(idxStr)

		if charSet[idxInt]{
			return false
		}
		charSet[idxInt]=true
	}
	return true
}


func main() {
	str := "sjdfoi"
	fmt.Println("-------------")
	start_1 := time.Now()
	is_1 := isUniqChars2(str)
	end_1 := time.Now()
	fmt.Printf("%t %f秒\n",is_1,(end_1.Sub(start_1)).Seconds())
	start_2 := time.Now()
	is_2 := isUniqChars(str)
	end_2 := time.Now()
	fmt.Printf("%t %f秒\n",is_2,(end_2.Sub(start_2)).Seconds())
	fmt.Println("-------------")

	str1 := "fhasuhhuhuh"
	start_1 = time.Now()
	is_1 = isUniqChars2(str1)
	end_1 = time.Now()
	fmt.Printf("%t %f秒\n",is_1,(end_1.Sub(start_1)).Seconds())
	start_2 = time.Now()
	is_2 = isUniqChars(str1)
	end_2 = time.Now()
	fmt.Printf("%t %f秒\n",is_2,(end_2.Sub(start_2)).Seconds())
	fmt.Println("-------------")


	str2 := "fhあああhuhuh"
	fmt.Println("-------------")
	start_1 = time.Now()
	is_1 = isUniqChars2(str2)
	end_1 = time.Now()
	fmt.Printf("%t %f秒\n",is_1,(end_1.Sub(start_1)).Seconds())
	start_2 = time.Now()
	is_2 = isUniqChars(str2)
	end_2 = time.Now()
	fmt.Printf("%t %f秒\n",is_2,(end_2.Sub(start_2)).Seconds())
	fmt.Println("-------------")
}
