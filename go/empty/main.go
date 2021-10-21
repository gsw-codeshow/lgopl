package main

import "fmt"

type Pair struct {
	V int
	D int
}

func main() {
	q := []Pair{{}}
	fmt.Println(q)
	fmt.Println(q[0].V, "---", q[0].D)
	fmt.Println("---", q[1:])
	return
}
