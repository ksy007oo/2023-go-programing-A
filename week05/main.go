package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input Score : ")
	reader := bufio.NewReader(os.Stdin)
	//inputScore := reader.ReadString('\n')    //1 variable but reader.ReadString returns 2 values
	//inputScore, err := read.ReadString('\n') // err declared and not used
	inputScoreString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	inputScoreString = strings.TrimSpace(inputScoreString)      // 공란 제거 ( 띄어쓰기 제거 )
	inputScore, err := strconv.ParseFloat(inputScoreString, 32) // 실제 형변환
	if err != nil {
		log.Fatal(err)
	}

	if inputScore >= 90 {
		grade := "A grade!"
	} else {
		grade := " BCDE grade"
	}
	fmt.Println("You got " + grade)
}
