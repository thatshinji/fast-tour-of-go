# Go-Trip
### 仓库采用Go规范的单元测试代码用作示例
1. 单元测试文件必须以_test.go结尾命名
2. 单元测试方法名须为 Test 开头, `function TestXXX()`
3. 单元测试方法参数必须是(t *Testing.T), `func TestGo(t *testing.T){/*...*/}`

example:

```go
package hello_world_test

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Log("Hello World")
}
```

### 目录
1. [go程序组成](https://github.com/thatshinji/fast-tour-of-go/tree/main/src/program_composition)
2. [go变量声明](https://github.com/thatshinji/fast-tour-of-go/tree/main/src/declaration)
3. [数据类型](https://github.com/thatshinji/fast-tour-of-go/tree/main/src/data_type)
4. [defer](https://github.com/thatshinji/fast-tour-of-go/tree/main/src/deferred)