package main

import (
	"fmt"
	"sync"

	"github.com/bigBarrage/roomManager/register"
	"github.com/bigBarrage/roomManager/run"
)

var wg sync.WaitGroup

func main() {
	register.RegisterProcessMessageFunc(ProcessMessage)
	rs := run.Run()
	fmt.Println("返回结果：", rs)
	wg.Add(1)
	wg.Wait()
}
