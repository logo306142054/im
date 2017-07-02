/*
* created by Orz on 2017/6/11
 */
package model

import (
	"im/imclient/log"
	. "im/public"
	"strconv"
	"sync"
)

type LoginedUsers struct {
	users map[string]User
	lock  sync.Mutex
}

var currentID int = 10000

var loginedUsers LoginedUsers

func init() {
	loginedUsers.users = make(map[string]User, 0)
}

func GetLoginedUsers() *LoginedUsers {
	return &loginedUsers
}

func (lu *LoginedUsers) AddNewUser(user User) (ID string) {
	lu.lock.Lock()
	defer lu.lock.Unlock()

	if user.Name == "" {
		log.Log("user name is empty")
		return ""
	}

	if _, ok := lu.users[user.ID]; ok {
		log.Log("exist user,ID=" + user.ID + " name=" + user.Name)
		return ""
	}

	if user.ID == "" {
		user.ID = strconv.Itoa(currentID)
		currentID++
	}

	lu.users[user.ID] = user
	return user.ID
}

func (lu *LoginedUsers) DelUser(ID string) {
	lu.lock.Lock()
	defer lu.lock.Unlock()

	if lu.isExistUser(ID) {
		log.Log("user leave,ID=" + ID)
		delete(lu.users, ID)
		return
	}
}

func (lu *LoginedUsers) GetUser(ID string) *User {
	lu.lock.Lock()
	defer lu.lock.Unlock()

	if user, ok := lu.users[ID]; ok {
		return &user
	}
	return nil
}

func (lu *LoginedUsers) isExistUser(ID string) bool {
	if _, ok := lu.users[ID]; ok {
		return true
	}
	return false
}
