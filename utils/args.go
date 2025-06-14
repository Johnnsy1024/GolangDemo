package utils

import (
	"fmt"
	"os"
)

func PrintArgs() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ { // 从非程序本体开始打印
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	s = ""
	sep = " "
	for _, arg := range os.Args { // 打印所有参数
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
