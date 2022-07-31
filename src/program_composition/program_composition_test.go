/**
Go程序的基本组成内容
*/
// 包名
package program_composition_test

// 导入golang内部包或第三方包
import (
	"fmt"
	"testing"
)

// 全局变量
const globalVar = 1

// 类型，类型声明首字母大些代表此类型可以被其他包使用
type MyInt int

// 方法，类型声明首字母大些代表此类型可以被其他包使用
func (mi MyInt) Print() {
	fmt.Println(mi)
}

// 函数
func TestProgramComposition(t *testing.T) {
	t.Log("Go 程序基本组成部分")
}
