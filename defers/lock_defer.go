package defers

import "sync"

/*
锁用完必须 Unlock，如果函数中间 return 或者 panic，忘记解锁会**死锁**。
defer 几乎是锁的标准搭档
规范：**Lock 之后立刻写 defer Unlock**，形成固定配对，防止漏写。
*/

var mu sync.Mutex
var count int

func Add() {
	mu.Lock()
	defer mu.Unlock() // 函数退出前自动解锁

	count++
	// 中间哪怕发生panic，也会执行Unlock释放锁！
}
