package main

import "fmt"

func ppp() {
	fmt.Println("这是一个中介函数")
}

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
