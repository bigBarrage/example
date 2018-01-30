package main

type Message struct {
	Type     int    `json:"type"`
	RoomID   string `json:"room_id"`
	Message  string `json:"message"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Sendtime int64  `json:"sendtime"`
}
