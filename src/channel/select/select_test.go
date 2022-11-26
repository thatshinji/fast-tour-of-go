package _select

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func TestSelect(t *testing.T) {
	var c1, c2 = generator(), generator()
	for {
		// select 从channel中随机收数据，谁快先收谁
		select {
		case n := <-c1:
			fmt.Println("Received from c1", n)
		case n := <-c2:
			fmt.Println("Received from c2", n)
			//default:
			//	fmt.Println("No value received")
			//}
		}
	}
}
