package deferred

import (
	"fmt"
	"testing"
)

func foo() (a, b int) {
	fmt.Println("exec foo")
	// 函数 foo 被 deferred 后 return 语句会被丢弃
	return 1, 2
}

func bar() (x, y int) {
	defer foo()
	return 1, 2
}

func TestFooBar(t *testing.T) {
	fmt.Println("test foo bar")
	x, y := bar()
	fmt.Println("x:", x, "y:", y)
}
