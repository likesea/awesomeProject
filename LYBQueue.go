package main

import (
	"fmt"
	"sync"
)

type LYBQueue struct {
	cond sync.Cond
	queque []interface{}
	cap int
	head int
	tail int
	locker   sync.Locker
}
func NewQueue(capacity int) *LYBQueue{
	return &LYBQueue{
		cond:   sync.Cond{L: &sync.Mutex{}},
		//有个空位
		queque: make([]interface{},capacity+1),
		cap:    capacity,
		head:0,
		tail:0,
		locker: &sync.Mutex{},
	}
}

func (p *LYBQueue) Pop()(interface{}) {
	//队列为空
	if p.head == p.tail {
		p.cond.L.Lock()
		fmt.Println("--空队列等待--")
		p.cond.Wait()
		//唤醒后解锁
		p.cond.L.Unlock()
	}
	//不为空，如果head<tail，队列形态如下
	// head    tail
	//  |       |
	//  - - - - -
	p.locker.Lock()
	defer p.locker.Unlock()
	//有新元素弹出去，广播
	defer p.cond.Broadcast()
	if p.head < p.tail {
		item := p.queque[p.head+1]
		p.head++
		return item
	}else {
		//不为空，如果head>tail，队列形态如下
		// tail    head
		//  |       |
		//  - - - - - -

		//head 指向最后一个位置
		if p.head == p.cap {
			//重置head位置到0
			p.head = 0
			item := p.queque[p.head ]
			return item
		} else {
			item := p.queque[p.head+1 ]
			p.head++
			return item
		}
	}
}


func (p *LYBQueue) Push(i interface{}) {
	p.locker.Lock()
	//队列为空
	if p.head == p.tail {
		//看是否位于数据最后
		if p.tail==p.cap{
			p.tail=0
			p.queque[p.tail] = i
		}else{
			p.tail++
			p.queque[p.tail] = i
		}
		//队列为空的时候，添加元素后广播
		p.cond.Broadcast()
		fmt.Println("--添加元素通知--")
	} else if p.head < p.tail {
	//不为空，如果head<tail，队列形态如下
	// head    tail
	//  |       |
	//  - - - - - -

		//队列满了，head在tail前面
		if p.tail==p.cap && p.head==0{
			p.cond.L.Lock()
			p.locker.Unlock()
			p.cond.Wait()
			p.cond.L.Unlock()
			p.locker.Lock()
			p.tail=0
			p.queque[p.tail] = i
			p.cond.Broadcast()
			//队列未满了，head在tail前面
		}else if p.tail==p.cap && p.head!=0{
			p.tail=0
			p.queque[p.tail] = i

		}else {
			p.tail++
			p.queque[p.tail] = i
		}
		p.cond.Broadcast()
		fmt.Println("--添加元素通知--")
	}else{
		//不为空，如果head>tail，队列形态如下
		// tail    head
		//          | |
		//  - - - - - -
		//队列满了，tail在head前面
		if  p.head==p.tail+1{
			p.cond.L.Lock()
			p.locker.Unlock()
			p.cond.Wait()
			p.cond.L.Unlock()
			p.locker.Lock()
			p.tail++
			p.queque[p.tail] = i
		}else {
			p.tail++
			p.queque[p.tail] = i
		}
		p.cond.Broadcast()
		fmt.Println("--添加元素通知--")
	}
	p.locker.Unlock()
}
