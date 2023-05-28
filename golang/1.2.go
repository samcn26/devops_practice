//模块一作业1.2
//基于 Channel 编写一个简单的单线程生产者消费者模型：
//
//队列：
//队列长度 10，队列元素类型为 int
//生产者：
//每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
//消费者：
//每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
package main

import (
	"fmt"
	"sync"
	"time"
)

// Queue 结构体用于表示队列
type Queue struct {
	queue []int      // 内部数组，用于存储元素
	lock  sync.Mutex // 互斥锁，用于保证线程安全
	cond  *sync.Cond // 条件变量，用于实现等待和通知的机制
}

// NewQueue 是 Queue 的构造函数，用于创建新的 Queue
func NewQueue(size int) *Queue {
	q := &Queue{}
	q.queue = make([]int, 0, size)
	q.cond = sync.NewCond(&q.lock)
	return q
}

// Enqueue 用于向队列中添加元素
func (q *Queue) Enqueue(item int) {
	q.lock.Lock() // 先获取锁
	for len(q.queue) >= cap(q.queue) {
		q.cond.Wait() // 如果队列满了，就等待
	}
	q.queue = append(q.queue, item)
	q.lock.Unlock()    // 释放锁
	q.cond.Broadcast() // 通知其他等待的 goroutine
}

// Dequeue 用于从队列中取出元素
func (q *Queue) Dequeue() int {
	q.lock.Lock() // 先获取锁
	for len(q.queue) == 0 {
		q.cond.Wait() // 如果队列空了，就等待
	}
	item := q.queue[0]
	q.queue = q.queue[1:]
	q.lock.Unlock()    // 释放锁
	q.cond.Broadcast() // 通知其他等待的 goroutine
	return item
}

// Producer 是生产者，会每秒生成一个新的元素并添加到队列中
func Producer(q *Queue) {
	for i := 0; ; i++ {
		q.Enqueue(i)
		time.Sleep(1 * time.Second)
	}
}

// Consumer 是消费者，会每秒从队列中取出一个元素并打印
func Consumer(q *Queue) {
	for {
		item := q.Dequeue()
		fmt.Println(item)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	q := NewQueue(10) // 创建一个新的队列
	go Producer(q)    // 启动生产者 goroutine
	Consumer(q)       // 启动消费者 goroutine
}
