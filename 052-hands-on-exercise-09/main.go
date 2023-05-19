package main

import (
	"fmt"
)

var i int = 5

const j int = 6

func main() {

	k := 5
	l := 42
	m := 24
	fmt.Printf("the value of i is %v and the type is i is %T\n", i, i)
	fmt.Printf("the value of j is %v and the type is j is %T\n", j, j)
	fmt.Printf("the value of k is %v and the type is k is %T\n", k, k)
	fmt.Printf("the value of l is %v and the type is l is %T\n", l, l)
	fmt.Printf("the value of m is %v and the type is m is %T\n", m, m)
	fmt.Println(i, j)
}
