/*
* created by Orz on 2017/6/12
 */
package model

import (
	"encoding/json"
	"im/imserver/log"
	. "im/public"
	"net"
	"sync"
)

var LastSessionID int = 0
var lock sync.Mutex

type Session struct {
	sid       int
	ip        string
	user      *User
	chatRoom  *ChatRoom
	sendCh    chan *CmdInfo
	leaveCh   chan int
	receiveCh chan []byte
	conn      net.Conn
}

func NewSession(newConn net.Conn) *Session {
	newsid := LastSessionID

	session := Session{
		sid:       newsid,
		ip:        newConn.RemoteAddr().String(),
		leaveCh:   make(chan int),
		receiveCh: make(chan []byte),
		sendCh:    make(chan *CmdInfo),
		conn:      newConn,
	}

	LastSessionID++

	return &session
}

func (s *Session) Ready() {
	go func() {
		for {
			select {
			case <-s.leaveCh:
				s.leave()
				return
			default:
			}

			select {
			case <-s.leaveCh:
				s.leave()
				return
			case ci := <-s.sendCh:
				pack := Package{
					CmdType: DATA,
					CmdInfo: *ci,
					Time:    GetCurrTime(),
				}
				b, err := json.Marshal(pack)
				if err != nil {
					log.Log("marshal pack failure")
					continue
				}
				s.conn.Write(Pack(b))
			}
		}
	}()
}

func (s *Session) SendMsg(cmdInfo *CmdInfo) {
	s.sendCh <- cmdInfo
}

func (s *Session) Close() {
	close(s.leaveCh)
}

func (s *Session) leave() {
	if s.user.ID == "" {
		return
	}

	pack := Package{
		CmdType: DATA,
		Time:    GetCurrTime(),
		CmdInfo: CmdInfo{
			Cmd:  LEAVE,
			Data: []byte(s.user.ID),
		},
	}

	Logout(&pack, s)
}
