package network

import (
	"encoding/json"
	. "im/imclient/model"
	. "im/public"
)

func HandleData(buff []byte, sendCh chan *CmdInfo) {
	pack, err := parsePackage(buff)
	if err != nil {
		return
	}

	switch pack.CmdType {
	case CONTROL, DATA:
		dispach(pack, sendCh)
	}
}

func parsePackage(buff []byte) (pack *Package, err error) {
	pack = new(Package)
	err = json.Unmarshal(buff, pack)
	return
}

type cbFunc func(pack *Package, sendCh chan *CmdInfo)

type stCBs struct {
	cmd int16
	cb  cbFunc
}

func dispach(pack *Package, sendCh chan *CmdInfo) {

	cbs := []stCBs{
		{LOGIN, RespLogin},
		{QUERY_ROOMS, RespQueryRooms},
		{ENTRY_CHATROOM, RespEntryRoom},
		{CHAT, RespBroadcast},
	}

	for _, cb := range cbs {
		if cb.cmd == pack.CmdInfo.Cmd {
			cb.cb(pack, sendCh)
		}
	}
}
