package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func returnRandomNumber() int {
	return rand.Intn(100)
}

type Callable func() int // 想在结构体中使用函数,则必须先定义这个函数类型

func ff() *int {
	v := 1
	return &v
}

func main() {
	switch returnRandomNumber() {
	case 0:
		println("Random number is 0")
	case 1:
		println("Random number is 1")
	case 2:
		println("Random number is 2")
	default:
		println("Random number is not 0, 1, or 2")
	}

	type myStruct struct {
		name   string
		age    int
		random Callable // Here we use the Callable type defined above
	}
	s := myStruct{}
	s.name = "John"
	fmt.Println(s.age)
	s.random = returnRandomNumber // Assign the function to the field
	fmt.Println(s.random())

	var f []string = []string{"apple", "banana", "orange"}
	fmt.Println(f)
	var p *int = new(int) // initialize a pointer to an int
	*p = 100

	a := 10
	hhh := &a
	*hhh = 20
	fmt.Println(a)

	o := 20
	l := 30
	if o >= 20 && l <= 30 {
		fmt.Println("我爱你")
	}
	fmt.Println(math.MaxFloat64)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from panic:", r)
		}
	}()
	// string_demo := "hello world"
	fmt.Println(strings.Join([]string{"hello", "world"}, ","))
	// fmt.Println(string_demo[0:20])
	string2 := "23"
	res, err := strconv.Atoi(string2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
