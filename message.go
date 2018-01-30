package main

import (
	"fmt"

	"github.com/bigBarrage/roomManager/room"
	"github.com/bigBarrage/roomManager/system"
	json "github.com/json-iterator/go"
)

func ProcessMessage(msg []byte, cn *room.ClientNode) {
	fmt.Println("你好")
	m := Message{}
	json.Unmarshal(msg, &m)
	fmt.Print("解析结果：")
	fmt.Printf("%+v\n", m)
	switch m.Type {
	case system.NODE_MESSAGE_TYPE_IN_HALL:
		cn.ChangeRoom("")
	case system.NODE_MESSAGE_TYPE_CHANGE_ROOM:
		fmt.Println("更换房间")
		cn.ChangeRoom(m.RoomID)
	default:
		fmt.Println("发送message")
		cn.SendMessageToRoom(m)
	}
}
