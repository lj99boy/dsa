package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func traceMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func f() {
	container := make([]int, 8)
	log.Println("> loop.")
	// slice会动态扩容，用它来做堆内存的申请
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
		if i == 16*1000*1000 {
			traceMemStats()
		}
	}
	log.Println("< loop.")
	// container在f函数执行完毕后不再使用
}

//执行的时候可以用GODEBUG=gctrace=1，显示当时内存分配和调用
//https://golang.org/pkg/runtime/
// runtime.GC()可以强制执行gc操作
// gc对内存使用的处理是，gc后无用的堆内存都被垃圾回收器回收了，但是不是立刻把内存还给操作系统，具体归还情况
// 当gctrace=>1的时候，当garbage collector归还内存时还会打印一条scavenging汇总信息
// 一个检测当前内存使用情况的是方法是 runtime.ReadMemStats()
func main() {
	log.Println("start.")
	f()

	log.Println("force gc.")
	runtime.GC() // 调用强制gc函数

	traceMemStats()
	log.Println("done.")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,syscall.SIGINT,syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sigChan)
}
