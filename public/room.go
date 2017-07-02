/*
* created by Orz on 2017/6/11
 */
package public

import "strconv"

type Room struct {
	ID      int    `json:"roomID"`
	Name    string `json:"roomName"`
	Info    string `json:"info"`
	Creator string `json:"creator"`
}

func NewRoom(name, info, creator string) *Room {
	room := &Room{
		Name:    name,
		Info:    info,
		Creator: creator,
	}
	return room
}

func (r *Room) String() string {
	str := "ID=" + strconv.Itoa(r.ID) + " Name=" + r.Name +
		" (" + r.Info + ") Creator=" + r.Creator

	return str
}
