package main

import "fmt"

var a string = "hello"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
)

func Slice(arrLens int) {
	// 切片的创建方式:必须带花括号或者用make
	arr := make([]int, arrLens)
	arr2 := []int{}
	arr3 := []int{1, 2, 3}
	bigSlice := [][]int{arr, arr2, arr3}

	for _, v := range bigSlice {
		fmt.Println(v)
	}
}

func Array() [5]int {
	// 数组的创建方式
	// arr := [...]int{1, 2, 3}
	arr2 := [5]int{1, 2, 3, 4, 5}
	return arr2
	// bigArr := [3][5]int{arr, arr2}
}

func main() {
	fmt.Println(a)
	fmt.Println(USD)
	fmt.Println(EUR)
	fmt.Println(GBP)
	// fmt.Println(Array())
}
