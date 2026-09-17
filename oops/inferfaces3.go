package oops

import (
	"fmt"
)

/*
空接口：没有定义任何方法，所有类型都自动实现空接口
`any` 只是 `interface{}` 的别名，二者完全等价。


类型断言：从接口变量取出底层值
语法：
value, ok := 接口变量.(具体类型)
- ok=true：类型匹配成功，value 是原始值
- ok=false：类型不匹配，不会 panic

*/

// 空接口
//var a interface{}

func funcName(a interface{}) string {
	value, ok := a.(string)
	if !ok {
		fmt.Println("It is not ok for type string")
		return ""
	}
	fmt.Println("The value is ", value)
	return value
}

// Go1.18+ any表示空接口
func printAnything(x any) {
	fmt.Println(x)
}

type IF interface{ foo() }
type S struct{}

func (s *S) foo() { fmt.Println("foo") }

func TestFaceNull() {
	var a int = 10
	funcName(a)

	printAnything(100)
	printAnything("hello")
	printAnything([]int{1, 2, 3})
	printAnything(map[string]int{"a": 1})

	var i IF
	i = &S{}
	i.foo()
}
