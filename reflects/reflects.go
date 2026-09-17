package reflects

/*
> 一句话核心：反射允许程序在运行时，获取变量的类型、字段、方法，动态读取 / 修改值，动态调用方法。
> Go 反射依托接口 interface**实现，只有存入接口的值，才能拿到反射信息。
> 典型场景：JSON 序列化、GORM ORM、配置解析、依赖注入。

⚠️ 反射是运行时能力，编译期无法做类型检查，代码可读性差、性能低，业务代码尽量少用
*/

import (
	"fmt"
	"reflect"
)

type User struct{}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestReflect() {

	var a int = 10
	// reflect.TypeOf()：拿到类型信息（结构体名、字段名、方法、tag）
	// reflect.ValueOf()：拿到值信息（可以读取、修改里面的数据）
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	//Type：变量具体类型（比如User自定义结构体）
	//Kind：底层基础类别，`struct / int / string / slice / map / ptr`
	fmt.Println("类型：", t)          // int
	fmt.Println("值：", v)           // 10
	fmt.Println("Kind：", t.Kind()) // int

	u := User{}
	t1 := reflect.TypeOf(u)
	fmt.Println(t1)        //reflects.User
	fmt.Println(t1.Kind()) // struct

	// 获取结构体信息（最常用场景，JSON/GORM 读取 struct tag）
	p := Person{"zhangsan", 18}
	t2 := reflect.TypeOf(p)

	fmt.Println("结构体名称：", t2.Name()) // Person

	fmt.Println(t2.Kind())

	for i := 0; i < t2.NumField(); i++ {
		f := t2.Field(i)
		fmt.Println(f.Tag.Get("json"))
		fmt.Println(f.Name)
		fmt.Println(f.Type)
	}

}
