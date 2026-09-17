package oops

import "fmt"

/*
核心一句话：**接口是一组方法的集合；Go 是隐式实现，
不需要写 `implements`。只要类型拥有接口全部方法，就自动实现这个接口**。
接口的价值：**定义行为，而不是定义类型，用来实现多态**

type 接口名 interface {
    方法1(参数) 返回值
    方法2(参数) 返回值
}
*/

// Speaker 接口：规定只要有 Say() 方法，就属于 Speaker
type Speaker interface {
	Say() string
}

// Person1 类型，实现 Say()
type Person1 struct {
	Name string
}

func (p Person1) Say() string {
	return "我是" + p.Name
}

// Dog 类型，实现 Say()
type Dog struct{}

func (d Dog) Say() string {
	return "汪 汪 汪"
}

// 接收Speaker接口，任何实现这个接口的类型都可以传进来
func Speak(speaker Speaker) {
	fmt.Println(speaker.Say())
}

func TestInterface() {
	p := Person1{Name: "zhangsan"}
	fmt.Println(p.Say())

	var s Speaker
	s = Person1{Name: "lisi"}
	Speak(s)

	s = Dog{}
	Speak(s)
}
