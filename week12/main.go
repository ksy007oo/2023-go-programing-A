// package main

// import "fmt"

// func main() {
// 	// var s []int
// 	// s = make([]int, 3)

// 	// s := make([]int, 5) // 단축 연산자 make 함수 이용해 슬라이스 생성 후 메모리 할당, 제로 값 적용

// 	s := []int{0, 0, 0, 0, 0} // 슬라이스 리터럴 이용해 슬라이스 생성 및 메모리 할당, 초기화 진행
// 	s[4] = 99
// 	s[2] = 91

// 	for _, value := range s {
// 		fmt.Println(value)
// 	}

// 	copyS := s[1:4]
// 	for _, value := range copyS {
// 		fmt.Println(value)
// 	}

// 	test := [3]string{"inha", "go", "student"} // 배열 리터럴 이용해 test 배열 생성
// 	// testS := test[0:4] // invalid argument: index 4 out of bounds [0:4]
// 	testS := test[:2] // testS := test[0:2]
// 	testS2 := test[1:]
// 	testS2[0] = "python"
// 	fmt.Println(testS2)
// 	fmt.Println(testS, len((testS)))
// 	fmt.Println(test)
// }

package main

import "fmt"

func main() {
	a := make([]string, 4, 5) // (타입, 개수, 수용량)
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2]
	as[1] = "z"
	// c := append(a, "y")
	c := append(a, "y", "x") // capacity가 바뀐다. 5 -> 10

	fmt.Println(a, len(a), cap(a))
	fmt.Println(c, len(c), cap(c))
	fmt.Printf("%x %x %x\n", &a[0], &as[0], &c[0])
	c[0] = "k"
	fmt.Println(a, len(a), cap(a))
	fmt.Println(a, c)
	// a := []string{"a", "b", "c", "d"}
	// as := a[:2]
	// as[1] = "z"
	// fmt.Println(a) // 원본도 바뀜
	// fmt.Println(as)

	// b := [4]int{4, 3, 2, 1}
	// bs := b[1:3]
	// fmt.Println(bs)

}
