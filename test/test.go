package main

import "fmt"

func main() {
	x, y := 10, 15

	fmt.Println(test(x, y))

	i := 1
	j := &i
	*j = 2

	var n *int
	m := 10
	n = &m
	u := n
	*n = 11
	*u = 12
	fmt.Println(n, m, u)

	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

func test(x, y int) int {

	return x + y
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}
