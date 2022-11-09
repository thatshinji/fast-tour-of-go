package deferred

import (
	"fmt"
	"testing"
)

func genReturns(i, j int) (x, y int) {
	defer func() {
		x = x * 5
		y = y * 10
	}()
	x = i + 1
	y = j + 2
	return
}

func TestReturns(t *testing.T) {
	x, y := genReturns(1, 2)
	fmt.Println("x: ", x, "y: ", y)
}

// output
//x:  10 y:  40
//--- PASS: TestReturns (0.00s)
