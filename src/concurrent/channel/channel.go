package channel

import "fmt"

var global int = 0
var c = make(chan int, 1)

func thread1() {
	fmt.Println("In thread 1")
	<-c
	global = 1
	c <- 1
}

func thread2() {
	fmt.Println("In thread 2")

	<-c
	global = 2
	c <- 1
}

func Run() {

	c <- 1 // Put the initial value into the channel
	go thread1()
	go thread2()

	fmt.Println("")
	fmt.Println("all done")
}
