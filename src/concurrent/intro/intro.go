package intro

import (
	"fmt"
	"math/rand"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	}
}
func Run() {

	go func(msg string) {
		f(msg)
	}("going-inline")

	go f("goroutine")

	f("direct")

	fmt.Println("")
	fmt.Println("all done")
}
