package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func newInt() *int {
	return new(int)
}

func main() {
	v := 1
	incr(&v)
	println(v)

	var p *int = new(int)
	*p = 1
	fmt.Println(p)

	k := newInt()
	*k = 2
	fmt.Println(*k)

}
