package main

import (
	"GolangDemo/sortAlgorithm"
	"GolangDemo/utils"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(utils.Add(1, 2))
	arr := []int{5, 2, 4, 6, 1, 3}
	fmt.Printf("Before Sort, arr is %v\n", arr)
	sortAlgorithm.BubbleSort(arr)
	ppp()
	fmt.Printf("After Sort, arr is %v\n", arr)
	logrus.Info("This is a log message")
}
