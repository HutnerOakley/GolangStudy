package test_var

import (
	"fmt"
)

const b string = "abc"
const c = "abc"

func Cons() {
	const LENGTH int = 10
	const WIDTH int = 20
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("area:%d\n", area)
	println(a, b, c)
}
