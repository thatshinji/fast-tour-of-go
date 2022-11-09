package declaration_test

// 只包含基本数据类型的声明，不包含slice,array, map等
import "testing"

// 因式分解关键字创建变量一般多用于声明全局变量
// const 关键字声明常量
const (
	constInt  = 1
	constBool = true
)

var (
	varInt  int64
	varBool bool
)

func TestDeclaration(t *testing.T) {
	// var 变量名 类型 =  变量值
	var a int64 = 1
	// 无显示类型声明，自动推导类型
	var b = 2
	// 短声明格式，采用 := 的方式
	c := 3
	// 同时声明同一类型多个字段
	var d, e, f int64 = 4, 5, 6
	var x, y = "hello", 3
	// 声明但不赋值，默认会使用零值
	var g int
	t.Log("a", a)
	t.Log("b", b)
	t.Log("c", c)
	t.Log("d", d)
	t.Log("e", e)
	t.Log("f", f)
	t.Log("g", g)
	t.Log("x", x)
	t.Log("y", y)
	t.Log("contInt", constInt)
	t.Log("contBool", constBool)
	t.Log("varInt", varInt)
	t.Log("varBool", varBool)
}

// output:
//declaration_test.go:29: a 1
//declaration_test.go:30: b 2
//declaration_test.go:31: c 3
//declaration_test.go:32: d 4
//declaration_test.go:33: e 5
//declaration_test.go:34: f 6
//declaration_test.go:35: g 0
//declaration_test.go:36: x hello
//declaration_test.go:37: y 3
//declaration_test.go:38: contInt 1
//declaration_test.go:39: contBool true
//declaration_test.go:40: varInt 0
//declaration_test.go:41: varBool false
