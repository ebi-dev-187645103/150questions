package main

import "fmt"

// 値numのiビット目がtrueであるかを返す
// Ex 1011 0,1,2ビット目→true, 3ビット目→false
// 1をiビットシフトし、0010000のような値を作る。
// 次にnumとのANDをとり、iビット目以外を0でクリアします。
// 最後にそれを0と比較し、0でなければもっとの値のiビット目は1、そうでなければ0
func GetBit(num,i int)bool{
	return (num & (1 << i)) != 0
}

// 値numにiビット目=trueをした値を返す
// まず、1をiビットシフトした00010000のような値を作る。
// 次にnumとのORをとり、iビット目だけ変更します。
// 他のビットは0でマスクされるので影響はありません。
func SetBit(num,i int)int{
	return num | (1 << i)
}

// このメソッドはSetBitとほぼ逆の操作。
// 値numにiビット目=falseをした値を返す
// まず、11101111のようなiビット目だけ0で残りが1になるような値を作る。
// これはnot演算で00010000の角ビットを反転させれば簡単にできます。
// 次にnumとのANDをとります。
// これでi番目のビットだけ0でクリアし、他のビットはそのままにしておくことができる。
func ClearBit(num,i int)int{
	mask := ^(1 << i)
	return num & mask
}

// 最上位ビットからiビット目までをクリアする
func ClearBitsMSBthroughI(num,i int)int{
	mask := (1 << i)-1
	return num & mask
}

// iビット目から0ビット目までをクリアする
func clearBitsIthrough0(num,i int)int{
	mask := ^((1 << (i+1)) -1)
	return num & mask
}

// このメソッドはSetBitとClearBitの考え方を合わせる。
// まず、ClearBitと同様に、11101111のようなマスクを使ってi番目のビットをクリアする。
// 次に置き換えたい値iビットをv(vは1または0)で置き換えることができる。
func UpdateBit(num,i,v int)int{
	mask := ^(1 << i)
	return (num & mask) | (v << i)
}

func main(){
	// GetBit
	b1 := 11 // 1011
	fmt.Println(GetBit(b1,1)) // 2:0010
	fmt.Println(GetBit(b1,2)) // 4:0100
	fmt.Println("--------------")
	// SetBit
	b2 := 11 // 1011
	fmt.Println(SetBit(b2,1)) // 1:0001 → 12:1100
	fmt.Println(SetBit(b2,2)) // 4:0100 → 15:1111
	fmt.Println("--------------")
	// ClearBit
	b3 := 11 // 1011
	fmt.Println(ClearBit(b3,1)) // 1:0001 → 9:1001
	fmt.Println(ClearBit(b3,2)) // 4:0100
	fmt.Println("--------------")
	// ClearBitsMSBthroughI
	b4 := 11 // 1011
	fmt.Println(ClearBitsMSBthroughI(b4,1)) // 1:0001 → 1:0001
	fmt.Println(ClearBitsMSBthroughI(b4,2)) // 4:0100 → 3:0011
	fmt.Println("--------------")
	// clearBitsIthrough0
	b5 := 27 // 11011
	fmt.Println(clearBitsIthrough0(b5,1)) // 1:00001 → 24:11000
	fmt.Println(clearBitsIthrough0(b5,3)) // 8:01000 → 16:10000
	fmt.Println("--------------")
	// UpdateBit
	b6 := 11 // 1011
	fmt.Println(UpdateBit(b6,1,2)) // 1:0001 → 13: 1101
	fmt.Println(UpdateBit(b6,3,2)) // 4:0100 → 19:10011
	fmt.Println("--------------")
}
