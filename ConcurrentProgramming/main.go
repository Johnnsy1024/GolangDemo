package main

import (
	"fmt"
	"sync"
)

func worker1(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("我是worker1,我将消息修改为100")
	ch <- "100"
}

func worker2(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("我是worker2,我将消息修改为200")
	ch <- "200"
}

func main() {
	ch := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go worker1(ch, wg)
	go worker2(ch, wg)
	value1 := <-ch
	value2 := <-ch
	fmt.Println("worker1修改后的消息:", value1)
	fmt.Println("worker2修改后的消息:", value2)
	wg.Wait()
	fmt.Println("执行完毕")
}
