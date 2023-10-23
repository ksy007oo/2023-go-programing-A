// 입력된 소수 판정 프로그램 v0.7
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"

// 	//"math/rand"
// 	"strconv"
// 	"strings"
// 	//"time" // seed 생성용 패키지
// )

// // 입력된 소수 판정 프로그램

// func main() {

// 	fmt.Print("정수 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n')

// 	if err != nil { // 에러가 발생하면
// 		log.Fatal(err)
// 	}

// 	in = strings.TrimSpace(in)
// 	number, err := strconv.Atoi(in)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true
// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break // 첫 번째 약수 발견되면 반복문 즉시 종료
// 		}
// 	}

// 	if isPrime { //비교 연산자 제거
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아입니다!")

// 	}
// }

// // 입력된 수의 소수 판정 프로그램 v0.8
// package main

// import (
// 	"fmt"
// 	"log"
// )

// // 입력된(fmt 패키지의 Scanln()) 소수 판정 프로그램

// func main() {
// 	var number int
// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&number)

// 	//fmt.Println(n)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true
// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break // 첫 번째 약수 발견되면 반복문 즉시 종료
// 		}
// 	}

// 	if isPrime { //비교 연산자 제거
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아입니다!")

// 	}
// }

// // 입력(0과 1처리)된 수의 소수 판정 프로그램 v0.9
// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// // 입력된(fmt 패키지의 Scanln()) 소수 판정 프로그램

// func main() {
// 	var number int

// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&number)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for number < 2 {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 		os.Exit(0)
// 	}

// 	isPrime := true
// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break // 첫 번째 약수 발견되면 반복문 즉시 종료
// 		}
// 	}

// 	if isPrime { //비교 연산자 제거
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아입니다!")

// 	}
// }

// 입력(0과 1처리)된 수의 소수 판정 프로그램 v0.9
package main

import (
	"fmt"
	"log"
	"os"
)

// 입력된(fmt 패키지의 Scanln()) 소수 판정 프로그램

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

	isPrime := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break // 첫 번째 약수 발견되면 반복문 즉시 종료
		}
	}

	if isPrime { //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아입니다!")

	}
}
