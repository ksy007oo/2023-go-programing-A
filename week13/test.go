package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args) // Args == Arguments (인수)
	fmt.Println(os.Args[2])

}
