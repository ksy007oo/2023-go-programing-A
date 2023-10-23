// package main

// import "fmt"

// func double(n *int) {
// 	*n = *n * 2
// }
// func main() {
// 	var a int = 6
// 	double(&a)
// 	fmt.Println(a)
// }

package main

import "fmt"

func swap(n1 *int, n2 *int) {
	//fmt.Println(n1, n2)
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

func main() {
	a := 10
	var b int = 20
	var pa *int = &a

	fmt.Printf("%T %T\n", &a, pa)
	fmt.Printf("%x %x %x\n", &a, pa, &pa)
	fmt.Println(&a, pa, &pa)
	fmt.Println(*pa)
	pa = &b
	fmt.Println(*pa)
}
