package main

import (
	"fmt"
	"reflect"

	//	"reflect"
	"strings"
)

func main() {
	texts := "G@ M@ney~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)

	//변수 명은 영문자로 시작해야함
	//영문 대문자인 경우 다른 패키지에서 접근할 수 있다. 소문자로 시작하는 변수는 동일 패키지에서만 접근 가능
	var e string
	var d bool
	var c rune // rune은 문자가아닌 숫자코드 저장.
	var b float64
	var a int

	// name convention
	var student_id string
	var i int32

	//a := 7
	fmt.Println(student_id)
	fmt.Println(i)

	// zero value
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))
}
