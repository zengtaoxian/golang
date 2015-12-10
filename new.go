package main

import "fmt"

func main() {
//	var p *[]int = new([]int)
//	var v []int = make([]int, 10)

	var p *[]int = new([]int)
	fmt.Println(p)

	*p = make([]int, 10, 10)
	fmt.Println(p)
	fmt.Println((*p)[2])

	v := make([]int, 10)
	fmt.Println(v)
}
