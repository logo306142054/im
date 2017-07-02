package model

import (
	"fmt"
	. "im/public"
	"sync"
)

type ServerInfo struct {
	chatroom *ChatRoom
	rooms    map[int]*Room
	look     sync.Mutex
}

var Server ServerInfo

func init() {
	Server.rooms = make(map[int]*Room)
}

func (si *ServerInfo) UpdateRooms(rooms []*Room) {
	si.look.Lock()
	defer si.look.Unlock()

	for _, r := range rooms {
		si.rooms[r.ID] = r
	}
}

func (si *ServerInfo) IsExitRoom(ID int) (*Room, bool) {
	si.look.Lock()
	defer si.look.Unlock()

	if _, ok := si.rooms[ID]; ok {
		return si.rooms[ID], true
	}
	return nil, false
}

func (si *ServerInfo) PrintAllRooms() {
	si.look.Lock()
	defer si.look.Unlock()

	fmt.Println("server room infos:")
	for _, r := range si.rooms {
		fmt.Println(r)
	}
}
