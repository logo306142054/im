/*
* created by Orz on 2017/6/12
 */
package model

import (
	. "im/public"
	"sync"
)

type ChatRoom struct {
	room     *Room
	sessions map[int]*Session
	look     sync.Mutex
}

var chatrooms = make([]*ChatRoom, 0)
var LastRoomID int = 1
var globalLook sync.Mutex

func SaveChatRoom(ct *ChatRoom) {
	globalLook.Lock()
	defer globalLook.Unlock()

	chatrooms = append(chatrooms, ct)
}

func GetChatRoom(ID int) *ChatRoom {
	globalLook.Lock()
	defer globalLook.Unlock()

	for _, ct := range chatrooms {
		if ct.room.ID == ID {
			return ct
		}
	}
	return nil
}

func GetAllRooms() []*Room {
	globalLook.Lock()
	defer globalLook.Unlock()

	rooms := make([]*Room, 0)
	for _, ct := range chatrooms {
		rooms = append(rooms, ct.room)
	}
	return rooms
}

func NewChatRoom(name, info, creator string) *ChatRoom {
	globalLook.Lock()
	defer globalLook.Unlock()

	r := NewRoom(name, info, creator)
	r.ID = LastRoomID
	LastRoomID++

	ct := &ChatRoom{
		room:     r,
		sessions: make(map[int]*Session),
	}
	return ct
}

func (ct *ChatRoom) AddNewSession(roomID int, session *Session) *Room {
	ct.look.Lock()
	defer ct.look.Unlock()

	if ct := GetChatRoom(roomID); ct != nil {
		ct.sessions[session.sid] = session
		session.chatRoom = ct
		return ct.room
	}
	return nil
}

func (ct *ChatRoom) DelSession(sid int) {
	ct.look.Lock()
	defer ct.look.Unlock()

	for _, s := range ct.sessions {
		if s.sid == sid {
			delete(ct.sessions, sid)
		}
	}
}
