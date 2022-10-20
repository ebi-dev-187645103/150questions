package main

import (
	"bytes"
	"fmt"
	"time"
	"unicode/utf8"
)

func InsertEach(s string, t string) string {
    var b bytes.Buffer
    for _, r := range s {
        if b.Len() > 0 {
            b.WriteString(t)
        }
        b.WriteRune(r)
    }
    return b.String()
}

func InsertEach2(s string, t string) string {
	ls, lt := len(s), len(t)
	r := make([]byte, 0, (lt+1)*ls)
	for i, w := 0, 0; i < ls; i += w {
			_, w = utf8.DecodeRuneInString(s[i:])
			if len(r) > 0 {
					r = append(r, ([]byte(t))...)
			}
			r = append(r, s[i:i+w]...)
	}
	return string(r)
}


func main() {
	start := time.Now()
	fmt.Println(InsertEach("abcdefghijklmnopqrstuvwxyz", "A"))
	end := time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	start = time.Now()
	fmt.Println(InsertEach("あいおえおかきくけこさしすせそ", "～"))
	end = time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	fmt.Println("-------------------------")

	start = time.Now()
	fmt.Println(InsertEach2("abcdefghijklmnopqrstuvwxyz", "A"))
	end = time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	start = time.Now()
	fmt.Println(InsertEach2("あいおえおかきくけこさしすせそ", "～"))
	end = time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
}
