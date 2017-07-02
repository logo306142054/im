package network

import (
	"fmt"
	"im/imserver/model"
	"im/imserver/util"
	. "im/public"
	"net"
	"os"
	"time"
)

func OpenSocket() {
	util.WaitGroup.Add(1)
	go listen()
}

func listen() {
	defer util.WaitGroup.Done()

	tcpListener, err := net.Listen("tcp4", ":10137")
	checkErr(err)

	defer tcpListener.Close()

	for {
		if conn, err := tcpListener.Accept(); err != nil {
			continue
		} else {
			go handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(time.Duration(1) * time.Minute))

	ip := conn.RemoteAddr().String()
	fmt.Println("new client:" + ip)

	session := model.NewSession(conn)
	session.Ready()

	readCh := make(chan []byte)
	go readBuff(readCh, session)

	readBuffFromNet(readCh, conn)

	session.Close()
	close(readCh)
}

func readBuffFromNet(readCh chan []byte, conn net.Conn) {
	buff := make([]byte, 2048)
	remaining := make([]byte, 0)
	var err error
	var readCount int

	for {
		if readCount, err = conn.Read(buff); err != nil {
			break
		}

		conn.SetDeadline(time.Now().Add(time.Duration(5) * time.Minute))
		remaining = append(remaining, buff[:readCount]...)
		remaining, err = Unpack(remaining, readCh)

		if err != nil {
			fmt.Printf("unpack occur err:%s\n", err.Error())
			continue
		}

		buff = make([]byte, 2048)
	}
}

func readBuff(readCh chan []byte, session *model.Session) {
	for buff := range readCh {
		handleData(buff, session)
	}
}

func notifyDisconnection() {
	//pack := Package{
	//	CmdType:DATA,
	//	Time: GetCurrTime(),
	//	CmdInfo: CmdInfo{
	//		Cmd:LEAVE,
	//		Data: []byte(""),
	//	},
	//}

}

func checkErr(err error) {
	if err != nil {
		os.Exit(1)
	}
}
