package main

import "fmt"

func main() {
	// var games map[int]string
	// games = make(map[int]string

	// games := make(map[int]string)

	// map literal
	games := map[int]string{
		456: "성기훈",
		218: "박해수",
		67:  "강새벽",
		1:   "오일남",
		199: "알리",
		101: "아이오아이",
	}

	// append ( 추가 )
	// games[456] = "성기훈"
	// games[218] = "박해수"
	// games[067] = "강새벽"
	// games[001] = "오일남"
	// games[199] = "알리"
	// games[101] = "아이오아이"

	// fmt.Println(gaems[100])
	// name, ok := games[100]
	name, ok := games[101]
	fmt.Println(name, ok)

	for _, v := range games {
		fmt.Println(v)
	}
	games[101] = "장덕수" // update ( 업데이트 )
	delete(games, 199) // delete ( 삭제 )

	for k, v := range games {
		fmt.Println(k, v)
	}

}
