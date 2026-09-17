package gotoutines

/*
1. goroutine（Go 协程）
goroutine 是 Go 运行时管理的**轻量级线程**，不是操作系统原生线程。
- 启动成本极低：初始栈很小（2KB），栈自动扩容 / 缩容；
- OS 线程一般 MB 级别，一个进程可以轻松跑上万 goroutine；
- 关键字：`go`，放在函数调用前面就启动一个协程。
⚠重点坑：`main` 是主 goroutine，main 函数一旦 return，
程序直接结束，不会等其他 goroutine 跑完。
等待协程的方案：
`time.Sleep`（仅测试用）、`sync.WaitGroup`、channel 阻塞等待。

## 2. channel 通道
channel 是 goroutine 之间通信的管道，用来在多个协程之间传递数据。
channel 是类型绑定的：只能传指定类型的数据。
语法:
	创建通道，第二个参数是缓冲区大小
	ch := make(chan int)         无缓冲通道
    chBuf := make(chan int, 2)   带缓冲通道，缓冲区容量2

两种 channel:
1. 无缓冲 channel
发送操作 `ch <- 1` 会阻塞，直到另一个 goroutine 执行接收 `<-ch`；
接收阻塞，直到有人发送。同步阻塞。
发送方和接收方必须同时准备好，才能完成通信。

2. 带缓冲 channel
缓冲区满之前，发送不会阻塞；缓冲区不为空，接收不会阻塞。
缓冲区满，发送阻塞；缓冲区空，接收阻塞。

3. select

`select`专门用来监听多个 channel 的 IO 操作，哪个channel就绪就执行哪个case。
类似 switch，但 case 都是 channel 收发操作。

规则：

1. 多个 case 同时就绪，**随机选一个执行**；
2. 所有 case 都阻塞，select 整体阻塞；
3. 加 `default`：没有 channel 就绪，直接执行 default，不会阻塞；
4. `case` 里可以是发送 `ch <- val`，也可以是接收 `v := <-ch`
*/

import "fmt"

func TestGoroutine() {
	c := make(chan int)

	go func() {
		defer fmt.Println("子go程结束")

		fmt.Println("子go程正在运行....")
		c <- 666 // 666发送到c
	}()

	num := <-c // 从c中接收数据，并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main go程结束")

}
