package test_var

import (
	"fmt"
)

// 全局变量
var GA int = 300
var GB int = 400

func Test1() {
	fmt.Println(GA, GB)
}

func init() {
	var a int
	fmt.Println("a=", a)
	fmt.Printf("%T\n", a)

	var b int = 100
	fmt.Println("b=", b)
	fmt.Printf("%T\n", b)

	var bb string = "hello"
	fmt.Println("bb=", bb)

	var c = 200
	fmt.Println("c=", c)
	fmt.Printf("%T\n", c)

	// 不支持全局变量
	e := 100
	fmt.Println("e=", e)
	fmt.Printf("%T\n", e)

	fmt.Println(GA)
	fmt.Println(GB)

}
