package deferred

import (
	"fmt"
	"testing"
)

func genPanic() {
	fmt.Println("gen a panic")
	panic(-1)
}

func panicFunc() {
	defer func() {
		if recover() != nil {
			fmt.Println("catch panic")
		}
	}()
	genPanic()
}

func TestPanicFunc(t *testing.T) {
	fmt.Println("exec start")
	panicFunc()
	fmt.Println("exit normally")
}

// output:
//exec start
//gen a panic
//catch panic
//exit normally
//--- PASS: TestPanicFunc (0.00s)
