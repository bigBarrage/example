package main

import (
	"fmt"
	"sync"

	"github.com/bigBarrage/roomManager/register"
	"github.com/bigBarrage/roomManager/room"
	"github.com/bigBarrage/roomManager/run"
)

var wg sync.WaitGroup

func main() {
	register.RegisterProcessMessageFunc(ProcessMessage)
	re := room.CreateRoom("festival")
	fmt.Println("创建房间结果：", re)
	rs := run.Run()
	fmt.Println("返回结果：", rs)
	wg.Add(1)
	wg.Wait()
}
