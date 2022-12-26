package deadlocks

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	lock := sync.Mutex{}

	var val int32
	increment := func() { //deadlock solution using mutex
		for i := 0; i < 100000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}

	//incrementAtomic := func() { //deadlock solution using atomic
	//	for i := 0; i < 100000000; i++ {
	//		atomic.AddInt32(&val, 1)
	//	}
	//	wg.Done()
	//}

	go increment()
	go increment()

	wg.Wait()
	fmt.Println(val)
}
