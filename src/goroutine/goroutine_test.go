package goroutine

import (
	"fmt"
	"testing"
	"time"
)

// go语言开一个进程，上面会有一个调度器，调度器负责调度协程，协程可能放在1个线程里面，可能2个协程放在1个线程里面。
// 任何函数只需要加上 go 关键字，就可以把函数交给调度器运行
// 调度器会在合适的时间进行切换

// goroutine 可能产生切换的点
// I/O, select
// channel
// 等待锁
// 函数调用（有时）
// runtime.Gosched()
// 以上只能作为参考，不能保证切换，不能保证在其他地方不切换

// goroutine 是一种非抢占式的机制

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				// fmt.Printf IO操作
				fmt.Printf("hello from goroutin: %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
