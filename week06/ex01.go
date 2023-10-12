package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var char rune = '가' // rune은 문자가아닌 숫자코드 저장.
	//var a int16 = 7
	//var a = 7
	//a := 7
	a := 7
	var b float64 = 5.3
	a = int(b) //형변환 Casting
	d := false

	fmt.Println(d)
	fmt.Printf("%T\n", d)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(char)        // 유니코드 값 출력
	fmt.Printf("%c\n", char) // 한글자 출력
	fmt.Printf("%T\n", char) // rune은 결국 int32의 별명

	fmt.Println(math.Round(2.71))
	fmt.Println("Hello Go~")
	fmt.Println(strings.Title("go git github java")) // 첫글자 대문자로 바꿔주는 strings의 Title() 함수

}
