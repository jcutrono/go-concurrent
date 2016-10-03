package singleton

import (
	"fmt"
	"sync"
	"time"
)

type onlyOneOfMe struct {
	Time time.Time
}

var instance *onlyOneOfMe
var mu sync.Mutex

func GetInstance() *onlyOneOfMe {

	if instance == nil {
		instance = &onlyOneOfMe{
			Time: time.Now(),
		}

		fmt.Println("I should only get created once!")
	}

	return instance
}

func Run() {

	//mu.Lock()
	//defer mu.Unlock()

	var wg sync.WaitGroup
	wg.Add(3)

	var sw sync.WaitGroup
	sw.Add(1)

	instances := make([]interface{}, 3)

	for i := 0; i < 3; i++ {

		go func(index int, sw *sync.WaitGroup, wg *sync.WaitGroup) {
			defer wg.Done()

			fmt.Println("waiting")
			sw.Wait()

			instances[index] = GetInstance()

		}(i, &sw, &wg)
	}

	time.Sleep(time.Second * 1)
	sw.Done()
	wg.Wait()

	for _, val := range instances {
		fmt.Println(val)
	}
}
