// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" // seed 생성용 패키지
// )

// func main() {
// 	// seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	for i := 1; i < 6; i++ {
// 		dice := rand.Intn(6) + 1
// 		fmt.Println(dice)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" // seed 생성용 패키지
// )

// // 난수 추출된 소수 판정 프로그램
// // 소수 : 1과 자기자신 외에는 나누어 떨어지지 않는 수(0과 1은 제외)
// func main() {
// 	// seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	count := 0
// 	number := rand.Intn(150) + 2 // 0과 1 제외. 2 ~ 151 사이의 수
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 1; i <= number; i++ {
// 		if number%i == 0 {
// 			count++
// 		}
// 	}

// 	if count == 2 {
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아입니다!")

// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" // seed 생성용 패키지
// )

// // 난수 추출된 소수 판정 프로그램
// // 소수 : 1과 자기자신 외에는 나누어 떨어지지 않는 수(0과 1은 제외)
// func main() {
// 	// seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	count := 0
// 	number := rand.Intn(150) + 2 // 0과 1 제외. 2 ~ 151 사이의 수
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i <= number; i++ { // 1과 nuber일 때 루프 돌지 않음
// 		if number%i == 0 {
// 			//count++
// 			count = count + 1
// 		}
// 	}

// 	if count == 0 {
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아입니다!")

// 	}
// }

// 난수 추출 프로그램 v0.3
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" // seed 생성용 패키지
// )

// // 난수 추출된 소수 판정 프로그램

// func main() {
// 	// seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	isPrime := true
// 	number := rand.Intn(150) + 2
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i <= number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			//count = count + 1
// 		}
// 	}

// 	if isPrime == true {
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아입니다!")

// 	}
// }

// // 난수 추출된 소수 판정 프로그램 v 0.4
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" // seed 생성용 패키지
// )

// // 난수 추출된 소수 판정 프로그램

// func main() {
// 	// seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	isPrime := true
// 	number := rand.Intn(150) + 2
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i <= number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			//count = count + 1
// 		}
// 		fmt.Print(i, " ")
// 	}

// 	if isPrime { //비교 연산자 제거
// 		fmt.Println(number, "는(은) 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아입니다!")

// 	}
// }

// 난수 추출된 소수 판정 프로그램 v 0.5
package main

import (
	"fmt"
	"math/rand"
	"time" // seed 생성용 패키지
)

// 난수 추출된 소수 판정 프로그램

func main() {
	// seed 설정
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	//number = 21
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i <= number; i++ {
		if number%i == 0 {
			isPrime = false
			break // 첫 번째 약수 발견되면 반복문 즉시 종료
		}
		//fmt.Print(i, " ")
	}

	if isPrime { //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아입니다!")

	}
}
