package channel

import (
	"fmt"
	"sync"
)

var global int = 0
var c = make(chan int, 1)

func thread1(wg *sync.WaitGroup) {

	msg := <-c // this blocks the channel
	fmt.Println("In thread 1: ", msg)

	setGlobal(msg)
	msg++
	c <- msg

	wg.Done()
}

func thread2(wg *sync.WaitGroup) {

	msg := <-c // this blocks the channel
	fmt.Println("In thread 2: ", msg)

	setGlobal(msg)
	msg++
	c <- msg

	wg.Done()
}

func setGlobal(i int) {

	fmt.Println("update global: ", i)
	global = i
}

func Run() {

	var wg sync.WaitGroup
	wg.Add(2)

	c <- 1 // Put the initial value into the channel
	go thread1(&wg)
	go thread2(&wg)

	wg.Wait()

	fmt.Println("")
	fmt.Println("all done")
}
