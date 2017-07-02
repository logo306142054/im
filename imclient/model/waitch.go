package model

import (
	. "im/public"
	"time"
)

var waitCh = make(chan int)

func WaitResp() bool {
	select {
	case err := <-waitCh:
		if err == ERR_SUCCESS {
			return true
		}
	case <-time.After(time.Duration(1) * time.Second):
		return false
	}
	return false
}
