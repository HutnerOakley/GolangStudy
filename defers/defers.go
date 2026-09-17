package defers

import (
	"fmt"
)

func Demo() string {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")

	return "demo..."

}
