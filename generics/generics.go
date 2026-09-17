package generics

/*
泛型:
让你可以写一套函数 / 结构体，支持多种类型，
编译期做类型检查，不用反射，性能和普通代码几乎一样。
解决痛点：
之前没有泛型时，相同逻辑要写多份函数（`IntMax`、`StringMax`）
或者被迫用`any`+ 反射（性能差、无编译校验）。

基础语法：类型参数
用 `[]` 放类型参数，`T` 是类型变量；`~`、`|`、约束用来限定允许什么类型。
*/

import (
	"fmt"
)

// T是类型参数；后面 int | string 是类型约束：T只能是int或string
func Max[T int | string](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

// 联合类型 `|` 多选
// `Number` 是接口作为类型约束，不是普通接口！泛型约束接口可以放类型集合。
type Number interface {
	int | int64 | float64
}

func Add[T Number](a, b T) T {
	return a + b
}

// `~` 底层类型
// `~T` 代表：底层类型是 T 的所有自定义类型
type MyInt int // MyInt底层类型是int
func TestMyInt[T ~int](x T) {
	fmt.Println(x)
}

func TestGeneric() {

	num1 := Max[int](10, 20)
	num2 := Max[string]("a", "b")
	// 类型推导：可以省略[int]/[string]，编译器自动推断
	num3 := Max(30, 40)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)

	num4 := Add(1, 2)
	fmt.Println(num4)

	a := 100
	TestMyInt(a) // ✅ 可以，MyInt底层是int

}
