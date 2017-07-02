package model

import (
	"bufio"
	"fmt"
	. "im/public"
	"os"
)

type ChatRoom struct {
	room        *Room
	sendCh      chan *CmdInfo
	broadcastCh chan string
	stopCh      chan int
}

func NewChatRoom(newRoom *Room, conn chan *CmdInfo) *ChatRoom {
	cr := ChatRoom{
		room:        newRoom,
		sendCh:      conn,
		broadcastCh: make(chan string),
		stopCh:      make(chan int),
	}
	return &cr
}

func (cr *ChatRoom) ReceiveMsg(msg string) {
	fmt.Printf("\n%s\n", msg)
}

func (cr *ChatRoom) Chat() {
	cmdInfo := CmdInfo{
		Cmd: CHAT,
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nyour say:")
	for scanner.Scan() {
		say := scanner.Text()
		if say == "#exit" {
			return
		}
		cmdInfo.Data = []byte(say)
		cr.sendCh <- &cmdInfo
	}
}
