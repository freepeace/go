package gcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

//关闭 receiving的channel
func TestCloseReceivingChan(t *testing.T){
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover();r != nil {
				t.Fatalf("CloseReceivingChan is panic")
			}
		}()
		wg.Done()
		t := <- c
		fmt.Println(t)
	}()
	wg.Wait() //确定goroutine 已经运行，这里不要使用 channel实现，这不是channel的正常功能，性能也不如WaitGroup
	time.Sleep(1)//确定goroutine 已经运行
	close(c)
}

//关闭 sending的channel
func TestCloseSendingChan(t *testing.T){
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover();r != nil {
				fmt.Println(r)
			}
		}()
		wg.Done()
		c <- 1
		t.Fatalf("CloseSendingChan is not panic")
	}()
	wg.Wait() //确定goroutine 已经运行，这里不要使用 channel实现，这不是channel的正常功能，性能也不如WaitGroup
	time.Sleep(1)//确定goroutine 已经运行
	close(c)
}

//recv的关闭chan
func TestRecvClosedChan(t *testing.T){
	c := make(chan int,1)
	c <- 1
	close(c)
	v,ok := <- c
	assert.Equal(t,1,v)
	assert.Equal(t,true,ok)
	v,ok = <- c
	assert.Equal(t,0,v)
	assert.Equal(t,false,ok)
}