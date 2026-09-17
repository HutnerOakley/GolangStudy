package slices

import (
	"fmt"
	"unsafe"
)

/*
slice：动态变长切片，
依托底层数组；语法 `[]T`，T 是元素类型（`[]int`、`[]string`）
核心要点：**不是数组，是视图；len 是当前元素数量，cap 是底层数组总容量**
*/

var a [3]int

func TestSlice() {
	fmt.Println(a)
	fmt.Println(a[len(a)-1])

	// 创建方式一：直接初始化，len=3, cap=3
	s := []int{10, 2, 3}
	fmt.Println(s, len(s), cap(s))

	// 创建方式二：make创建（最常用，可以指定 len、cap）
	// make([]T, len, cap)`，cap 可以省略，省略时 cap = len
	s1 := make([]int, 2, 5)
	fmt.Println(s1)

	// 省略cap：cap等于len
	s3 := make([]int, 3)
	fmt.Println(len(s3), cap(s3))

	// 从数组或者已有 slice 切片得到（视图，共享底层数组）
	arr := [5]int{1, 2, 3, 4, 5}
	s4 := arr[1:4] // 左闭右开：下标1、2、3 → [2,3,4]
	fmt.Println(s4)

	// 空切片 vs nil切片
	var s5 []int
	s6 := []int{}
	fmt.Println(s5)
	fmt.Println(s6)

	// 利用 unsafe 查看底层指针
	fmt.Println(*(*uintptr)(unsafe.Pointer(&s5))) // 输出0
	fmt.Println(*(*uintptr)(unsafe.Pointer(&s6))) // 输出一个非零内存地址

	fmt.Println(s5 == nil) // true
	fmt.Println(s4 == nil) // false

	// 访问元素&遍历

	s7 := []int{10, 20, 30}
	fmt.Println(s7)
	s7[1] = 200

	// for下标遍历
	for i := 0; i < len(s7); i++ {
		fmt.Println(s7[i])
	}

	// range遍历
	for idx, val := range s7 {
		fmt.Println(idx, val)
	}

	// append追加元素
	s8 := []int{1, 2}
	s8 = append(s8, 3)
	s8 = append(s8, 4, 5, 6)
	s9 := []int{7, 8}
	s8 = append(s8, s9...)
	fmt.Println(s8)

}
