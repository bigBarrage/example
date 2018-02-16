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
	/*
		opt := config.BroadcastingStationOption{}
		opt.Host = "127.0.0.1"
		opt.Port = 8911
		opt.Path = "broadcasting"

		config.SetBroadcastingStation(opt)
	*/
	register.RegisterProcessMessageFunc(ProcessMessage)
	//register.RegisterProcessMessageFromBroadcastingFunc(ProcessMessageFromBroadcasingStation)
	re := room.CreateRoom("festival")
	fmt.Println("创建房间结果：", re)
	rs := run.Run()
	fmt.Println("返回结果：", rs)
	wg.Add(1)
	wg.Wait()
}
