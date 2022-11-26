package done

import (
	"fmt"
	"sync"
	"testing"
)

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("id: %d, chan: %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in chan int
	//wg *sync.WaitGroup
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func channelDemo() {
	//var c chan int // c == nil
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	//for _, worker := range workers {
	//	<-worker.done
	//}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	//for _, worker := range workers {
	//	<-worker.done
	//}
	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
	wg.Wait()
}

func TestChannel(t *testing.T) {
	channelDemo()
}
