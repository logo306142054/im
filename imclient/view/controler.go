package view

import (
	"fmt"
	"im/imclient/model"
	. "im/public"
	"strconv"
)

func WaitForInput(sendCh chan *CmdInfo, stopCh chan int) {

	for {
		select {
		case <-stopCh:
			return
		default:
		}

		select {
		case <-stopCh:
			return
		default:
			if !input(sendCh) {
				return
			}
		}
	}
}

func input(sendCh chan *CmdInfo) bool {
LOOP:
	fmt.Println("***************************")
	fmt.Println("*   1、show all chatrooms *")
	fmt.Println("*   2、choose chatroom    *")
	fmt.Println("*   3、create chatroom    *")
	fmt.Println("*   4、exit               *")
	fmt.Println("***************************")
	fmt.Print("your selection:")

	var selection int

	fmt.Scanf("%d\n", &selection)

	switch selection {
	case 1:
		showChatRoom()
	case 2:
		entryChatRoom(sendCh)
	case 3:
	case 4:
		return false
	default:
		goto LOOP
	}
	return true
}

func showChatRoom() {
	model.Server.PrintAllRooms()
}

func entryChatRoom(sendCh chan *CmdInfo) {
	fmt.Print("input your choice room ID:")
	var ID int
	fmt.Scanf("%d\n", &ID)

	room, ok := model.Server.IsExitRoom(ID)
	if !ok {
		return
	}
	//ENTRY_CHATROOM
	cmdInfo := CmdInfo{
		Cmd:  ENTRY_CHATROOM,
		Data: []byte(strconv.Itoa(ID)),
	}

	sendCh <- &cmdInfo

	if !model.WaitResp() {
		return
	}

	model.CurChatRoom = model.NewChatRoom(room, sendCh)
	model.CurChatRoom.Chat()

}
