// 소수 판정 프로그램 v1.0 : 함수 적용, error 
package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) bool {
	prime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}
	return prime // true 리턴하면 소수, false 소수가 아님.
}

// 입력된(fmt 패키지의 Scanln()) 소수 판정 프로그램 : 함수 적용

func main() {
	var number int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number)

	if err != nil {
		log.Fatal(err)
	}

	for number < 2 {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
		os.Exit(0)
	}

	if isPrime(number) { //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아입니다!")

	}
}

// 소수 판정 프로그램 v0.9
package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) bool {
	prime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}
	return prime // true 리턴하면 소수, false 소수가 아님.
}

// 입력된(fmt 패키지의 Scanln()) 소수 판정 프로그램 : 함수 적용

func main() {
	var number int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number)

	if err != nil {
		log.Fatal(err)
	}

	for number < 2 {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
		os.Exit(0)
	}

	if isPrime(number) { //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아입니다!")

	}
}

// // after(multi return)
// package main

// import (
// 	"fmt"
// )

// func processScore(name string, kor int, eng int, math int) (int, int) {
// 	totalScore := kor + eng + math
// 	average := totalScore / 3

// 	return totalScore, average
// 	//fmt.Printf("%s님의 총점은 %d, 평균은 %d점 입니다.\n", name, totalScore, average)
// }

// func main() {
// 	var t int
// 	var a int
// 	t, a = processScore(100, 90, 93)
// 	fmt.Printf("%s님의 총점은 %d, 평균은 %d점 입니다.\n", "홍길동", t, a)

// 	t, a = processScore(89, 91, 92)
// 	fmt.Printf("%s님의 총점은 %d, 평균은 %d점 입니다.\n", "홍길동", t, a)
// }

// // after
// package main

// import (
// 	"fmt"
// )

// func processScore(name string, kor int, eng int, math int) {
// 	totalScore := kor + eng + math
// 	average := totalScore / 3

// 	fmt.Printf("%s님의 총점은 %d, 평균은 %d점 입니다.\n", name, totalScore, average)
// }

// func main() {
// 	processScore("길동이", 100, 90, 93)
// 	processScore("길동짱 ", 100, 90, 93)

// kor = 100
// eng = 90
// math = 93
// name = "길동이"

// fmt.Println(name, "의 총점은", (kor + eng + math), "입니다. 평균은", (kor+eng+math)/3.0)

// kor = 100
// eng = 90
// math = 93
// name = "길동짱"

// fmt.Println(name, "의 총점은", (kor + eng + math), "입니다. 평균은", (kor+eng+math)/3.0)

//}

// // before
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	kor := 100
// 	eng := 90
// 	math := 93
// 	name := "길동이"

// 	fmt.Println(name, "의 총점은", (kor + eng + math), "입니다. 평균은", (kor+eng+math)/3.0)

// 	kor = 100
// 	eng = 90
// 	math = 93
// 	name = "길동짱"

// 	fmt.Println(name, "의 총점은", (kor + eng + math), "입니다. 평균은", (kor + eng + math)/3.0)

// }

// package main

// import (
// 	"fmt"
// )

// func sayHello() {
// 	fmt.Println("Hello~")
// }

//	func main() {
//		sayHello()
//	}
//
