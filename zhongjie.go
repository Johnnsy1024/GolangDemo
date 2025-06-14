package main

import "fmt"

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
