package main

import (
	"fmt"
	"sync"
)

// "fmt"

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
	// s := []int{1, 2, 3}
	// utils.ChangeSlice(s)
	// fmt.Printf("Reference modified array s: %v\n", s)
	// utils.ChangeSlice2(&s)
	// fmt.Printf("Pointer modified array s: %v\n", s)
	// rand.Seed(time.Now().UnixNano())
	// utils.Lissajous(os.Stdout)
	// utils.GetWebUrl()
	// go utils.Hello()
	// fmt.Println("Hello from main()")
	// time.Sleep(time.Second * 5)

	// 标准的并发模式
	// wg := &sync.WaitGroup{}
	// for i := 1; i <= 5; i++ {
	// 	wg.Add(1)
	// 	go worker(i, wg)
	// }
	// wg.Wait()
	// fmt.Println("All workers finished")
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan string)
	go worker1(ch, wg)
	go worker2(ch, wg)

	value1 := <-ch
	value2 := <-ch
	wg.Wait()
	fmt.Println("Received message 1:", value1)
	fmt.Println("Received message 2:", value2)
}
