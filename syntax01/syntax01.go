// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var originalCount int = 10
// 	fmt.Println("I started with", originalCount, " apples.")
// 	var eatenCount int = 4
// 	fmt.Println("Some jerk ate", eatenCount, "apples.")
// 	fmt.Println("There are", originalCount-eatenCount, "apples left.")
// }

// ========================================================

package main

import (
	"fmt"
	"strings"
)

func main() {
	texts := "G@ M@ney~~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newtexts := r.Replace(texts)
	fmt.Println(newtexts)

	// 변수명은 영문자로 시작해야함'
	// 대문자의 경우 다른 패키지 접근 가능. 소문자 시작하는 변수는 동일 패키지에서만 가능
	// var e string
	// var d bool
	// var c rune
	// var b float64
	// var a int
	// //a := 7

	// // naming convention
	// var studentId string
	// var i int32

	// fmt.Println(studentId)
	// fmt.Println(i)

	// //zero value
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// fmt.Printf("%T\n", d)
	// fmt.Printf("%T\n", e)

	// fmt.Println(reflect.TypeOf(d))
	// fmt.Println(reflect.TypeOf(e))
}

// ========================================================
