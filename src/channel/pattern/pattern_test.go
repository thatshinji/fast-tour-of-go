package pattern

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 生成器模式
func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s message %d", name, i)
			i++
		}

	}()
	return c
}

// 生成器模式
// 当入参比较少且已知的时候可以使用该种模式
func fanIn(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()
	go func() {
		for {
			c <- <-c2
		}
	}()
	return c
}

func fanInByRestsParams(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		ch := ch // 注意变量赋值
		go func() {
			for {
				c <- <-ch
			}
		}()
	}
	return c
}

// 建议用此模式
// 且入参不确定时，更建议使用此模式
func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}

		}
	}()
	return c
}

func TestPattern2(t *testing.T) {
	m := msgGen("service1")
	m1 := msgGen("service2")
	//c := fanIn(m, m1)
	c := fanInBySelect(m, m1)
	for {
		fmt.Println(<-c)
	}
}

func TestPattern(t *testing.T) {
	m := msgGen("service1")
	m1 := msgGen("service2")
	for {
		fmt.Println(<-m)
		fmt.Println(<-m1)
	}
}
