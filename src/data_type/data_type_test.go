package data_type_test

import (
	"testing"
)

// 基本数据类型
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32, represents a Unicode code point
// float32 float64
// complex64 complex128 // 复数类型

// 类型别名
type MyInt int64

func TestDataType(t *testing.T) {
	var i int64 = 32
	var i2 MyInt = 32
	// 代码无法通过测试
	// errors: invalid operation: i == i2 (mismatched types int64 and MyInt)
	//t.Log("i isn't equals i2 :", i == i2) // 可取消代码注释尝试运行该行代码会报错
	// 显示类型转换
	// data_type_test.go:25: i is equals i2 true
	t.Log("i is equals i2", i == int64(i2))
}

func TestStringAppend(t *testing.T) {
	h := "hello" // string类型必须用双引号
	w := "world"
	t.Log(h + " " + w)
	// data_type_test.go:32: hello world
}

func TestStringTemplate(t *testing.T) {
	// 反引号不支持任何转义序列
	s := `hello\n
	go`
	t.Log(s)
	//data_type_test.go:39: hello\n
	//go
}

func TestRune(t *testing.T) {
	var c rune = '你' // 注意单引号
	t.Logf("c=%v ct=%T\n", c, c)
	//data_type_test.go:48: c=20320 ct=int32
}
