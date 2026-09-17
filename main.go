package main

import (
	"fmt"
	_ "go_project/InitLib1"
	_ "go_project/InitLib2"
	"go_project/defers"
	_ "go_project/defers"
	"go_project/generics"
	"go_project/gotoutines"
	"go_project/oops"
	"go_project/reflects"
	"go_project/servers"
	"go_project/slices"
)

func init() {
	fmt.Println("libmain init")
}

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	//s := "gopher"
	//fmt.Println("Hello and welcome, %s!", s)
	//
	//for i := 1; i <= 5; i++ {
	//	//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
	//	// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
	//	fmt.Println("i =", 100/i)
	//}
	//
	//test_var.Test1()
	//test_var.Cons()
	//x, y := funcs.Swap("Mahesh", "Kumar")
	//fmt.Println(x, y)
	//
	//a := 100
	//b := 200
	//
	//fmt.Printf("交换前，a 的值 : %d\n", a)
	//fmt.Printf("交换前，b 的值 : %d\n", b)
	//
	//funcs.Swaps(&a, &b)
	//
	//fmt.Printf("交换后，a 的值 : %d\n", a)
	//fmt.Printf("交换后，b 的值 : %d\n", b)
	//
	//demo := defers.Demo()
	//fmt.Println(demo)

	//recovers.Demo(10)
	////产生错误后 程序继续
	//fmt.Println("程序继续执行...")
	//
	//n, err := defers.ReadFile("D:\\go_project\\var")
	//fmt.Println(n, err)

	defers.Add()

	slices.TestSlice()
	slices.TestMap()

	oops.TestStruct()
	oops.TestUser()
	oops.TestInterface()
	oops.TestFaces()
	oops.TestFaceNull()

	reflects.TestReflect()

	generics.TestGeneric()

	gotoutines.TestGoroutine()

	// 启动server
	server := servers.NewServer("127.0.0.1", 8888)
	server.Start()

}
