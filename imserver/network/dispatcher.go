/*
* created by Orz on 2017/6/11
 */
package network

import (
	"encoding/json"
	"fmt"
	"im/imserver/model"
	. "im/imserver/model"
	. "im/public"
)

func handleData(buff []byte, session *model.Session) {
	if pack, err := parsePackage(buff); err != nil {
		fmt.Println("parse package failure, ", err.Error())
		return
	} else {
		switch pack.CmdType {
		case CONTROL, DATA:
			dispach(pack, session)
		}
	}
}

func parsePackage(buff []byte) (pack *Package, err error) {
	pack = new(Package)
	err = json.Unmarshal(buff, pack)
	return
}

type cbFunc func(pack *Package, session *model.Session)

type stCBs struct {
	cmd int16
	cb  cbFunc
}

func dispach(pack *Package, session *model.Session) {
	cbs := []stCBs{
		{LOGIN, Login},
		{LEAVE, Logout},
		{QUERY_ROOMS, QueryRooms},
		{CHAT, Chat},
		{ENTRY_CHATROOM, EntryRoom},
	}

	for _, cb := range cbs {
		if cb.cmd == pack.CmdInfo.Cmd {
			cb.cb(pack, session)
		}
	}
}
