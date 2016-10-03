package updateone

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func Run() {

	var unsafeInt uint64 = 0
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				unsafeInt++
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second * 1)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Printf("ops: %v | unsafeInt: %v\n", opsFinal, unsafeInt)
}
