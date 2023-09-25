package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Input Score : ")
	reader := bufio.NewReader(os.Stdin)
	//inputScore := reader.ReadString('\n')    //1 variable but reader.ReadString returns 2 values
	//inputScore, err := read.ReadString('\n') // err declared and not used
	inputScore, err := reader.ReadString('\n') //option 2
	log.Fatal(err)
	fmt.Println(inputScore)
}
