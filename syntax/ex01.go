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

package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var c rune = '가' // 문자열은 "" 룬은 ''
	//var a int16 = 7
	//var a = 7
	//a := 7
	a := 7
	var b float64 = 5.3
	a = int(b) // 캐스팅 , 형변환
	d := false

	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(c)        //44032
	fmt.Printf("%c\n", c) // 가
	fmt.Printf("%T\n", c) // int32 -> rune 은 결국 int32의 별명

	fmt.Println(math.Ceil(2.71))                     // 3
	fmt.Println("Hello go")                          // Hello go
	fmt.Println(strings.Title("go git github java")) // Go Git Github Java
}
