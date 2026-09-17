package oops

/*
结构体（`struct`）是**自定义复合类型**，
把多个不同类型的字段打包成一个整体，是Go实现面向对象的基础单元。
类比：Java 的 class、Python 的 dataclass。

type 结构体名 struct {
    字段名 类型
    字段名 类型
    // ...
}
*/

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type Person struct {
	Name string
}

func (p *Person) say() {
	fmt.Println("hello", p.Name)
}

type Student struct {
	Person
	Score int
}

type Teacher struct {
	P      Person
	Course string
}

type Class struct {
	Name string
	Age  int
}

func TestUser() {
	// 方式1：按顺序初始化（不推荐，字段顺序改了就炸）
	u1 := User{1, "zhangsan", 20}

	// 方式2：键值对初始化【推荐】，顺序无关
	u2 := User{
		ID:   1,
		Name: "zhangsan",
		Age:  20,
	}

	fmt.Println(u1.Name, u2.Age)

	// 访问结构体字段
	fmt.Println(u2.Name)
	u2.Age = 23 // 修改字段
	fmt.Println(u2)

	// 取结构体地址
	u3 := &User{ID: 3, Name: "wangwu", Age: 18}

	// Go语法糖：指针访问字段不用 -> ，直接 .
	fmt.Println(u3.Name)
	u3.Age = 19
	fmt.Println(u3)

	// 结构体变量声明不赋值时，所有字段**自动填充各自类型的零值
	var u4 User
	fmt.Println(u4)

	// 结构体嵌入（组合，用来替代继承）
	//匿名嵌入：不写字段名，直接写类型，方法会提升到外层结构体。
	//不是继承，Student 和 Person 是两种独立类型，不存在父子关系。
	s := Student{
		Person: Person{Name: "lisi"},
		Score:  20,
	}
	s.say() // 直接调用Person的方法，语法糖 s.Person.Hello()

	// 命名嵌入
	t := Teacher{
		P:      Person{Name: "wangwu"},
		Course: "yuwen",
	}
	t.P.say() // 需要使用t.P.Hello()，不能直接 t.Hello()

	c := Class{
		Name: "feng",
		Age:  18,
	}
	data, _ := json.Marshal(c)
	fmt.Println(string(data))

}

// 定义一个结构体
type T struct {
	name string
}

func (t T) test1() {
	t.name = " new test1"
}

func (t *T) test2() {
	t.name = "new test2"
}

func TestStruct() {

	t := T{"old name"}

	fmt.Println("test1 调用前", t.name)
	t.test1()
	fmt.Println("test1 调用后 ", t.name)

	fmt.Println("test2 调用前", t.name)
	t.test2()
	fmt.Println("test2 调用前", t.name)

}
