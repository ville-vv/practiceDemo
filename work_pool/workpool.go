package main

import (
	"time"
	"fmt"
)


type Task struct{
	f func() error
}

func NewTask(f func() error)*Task{
	t := &Task{
		f:f,
	}
	return t
}
func (t *Task)Execute(){
	t.f()
}


type Pool struct{
	EntryChannel chan *Task
	JobsChannel chan *Task
	worker_num int 
}

func NewPool(cap int)*Pool{
	p := Pool{
		EntryChannel:make(chan *Task),
		JobsChannel:make(chan *Task),
		worker_num:cap,
	}

	return &p

}

func(p *Pool)worker(worker_ID int){
	for task := range p.JobsChannel{
		task.Execute()
		fmt.Println("worker_ID", worker_ID, "执行完成")
	}
}


func (p *Pool)Run(){
	for i:=0; i< p.worker_num; i++{
		go p.worker(i)
	}

	for task := range p.EntryChannel{
		p.JobsChannel <- task
	}

	close(p.JobsChannel)
	close(p.EntryChannel)
	fmt.Println("程序结束")
}

func RunPool(){
	t := NewTask(func()error{
		fmt.Println(time.Now())
		return nil
	})
	t1 := NewTask(func()error{
		fmt.Println("task 2")
		return nil
	})
	p := NewPool(3)
	go func(){
		for{
			p.EntryChannel <-t
			p.EntryChannel <-t1
			time.Sleep(time.Second*1)
		}
	}()
	p.Run()
}


func ChanFor(){
	var ch chan int
	ch = make(chan int)
	go func(){
		for{
			ch <-1
			time.Sleep(time.Microsecond*1)
		}
		
	}()

	for c := range ch{
		fmt.Println("循环获取：", c)
	}
}


func main(){
	go ChanFor()
	 RunPool()

}