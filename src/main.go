package main

import (
	"concurrent/channel"
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(4)
	fmt.Println("Logical CPUs: ", runtime.NumCPU())

	//intro.Run()

	//updateone.Run()

	channel.Run()

	//singleton.Run()
}
