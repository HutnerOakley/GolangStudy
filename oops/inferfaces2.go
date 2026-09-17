package oops

import (
	"fmt"
)

/*
值接收者 vs 指针接收者（高频坑）

> 规则：
> 1. 如果接口方法是**值接收者 `func (p Person) Say()`
> `Person` 和 `*Person` 都实现接口（Go 自动解引用）
> 2. 如果接口方法是**指针接收者 `func (p *Person) Say()`
> 只有 `*Person` 实现接口，Person 值类型不满足！
*/

type Speaker1 interface {
	Say() string
}

type Person2 struct {
	Name string
}

// 指针接收者
func (p *Person2) Say() string {
	return p.Name
}

func TestFaces() {
	var s Speaker1 = &Person2{"wangwu"}
	//var s1 Speaker1 = Person2{"wangwu"} // ❌ 编译报错！Person值类型没有实现
	fmt.Println(s.Say())
}
