package main

import (
	"github.com/bigBarrage/roomManager/room"
	"github.com/bigBarrage/roomManager/system"
	json "github.com/json-iterator/go"
)

func ProcessMessage(msg []byte, cn *room.ClientNode) {
	m := Message{}
	json.Unmarshal(msg, &m)
	switch m.Type {
	case system.NODE_MESSAGE_TYPE_IN_HALL:
		cn.ChangeRoom("")
	case system.NODE_MESSAGE_TYPE_CLEAN_ROOM:
		cn.ChangeRoom(m.RoomID)
	default:
		cn.SendMessage(m)
	}
}
