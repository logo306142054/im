package network

import (
	"encoding/json"
	"fmt"
	"im/imclient/log"
	"im/imclient/util"
	. "im/public"
	"net"
	"os"
)

func Connect(sendCh chan *CmdInfo, stopCh chan int) {
	dial("127.0.0.1", "10137", sendCh, stopCh)
}

func dial(ip, port string, sendCh chan *CmdInfo, stopCh chan int) {
	conn, err := net.Dial("tcp", ip+":"+port)
	checkErr(err)

	util.WaitGroup.Add(1)
	go handleConn(conn, sendCh, stopCh)
}

func handleConn(conn net.Conn, sendCh chan *CmdInfo, stopCh chan int) {
	defer util.WaitGroup.Done()
	defer conn.Close()

	var readCh = make(chan []byte)

	go readBuff(readCh, sendCh)
	go sendBuff(sendCh, stopCh, conn)

	readBuffFromNet(readCh, conn)

	close(stopCh)
	fmt.Println("conn close")
}

func readBuffFromNet(readCh chan []byte, conn net.Conn) {
	var remaining = make([]byte, 0)
	var err error

	var buff = make([]byte, 2048)
	var totalRead int
	for {
		if totalRead, err = conn.Read(buff); err != nil {
			break
		}
		//将本次读取到的数据与上次未完成解析的数据拼接到要给切片中一起解析（解决粘包的问题）
		remaining = append(remaining, buff[:totalRead]...)
		//如果解析成功，且又多余的数据但又无法组成完整包的数据将赋值给remaining，否则返回nil
		remaining, err = Unpack(remaining, readCh)
		if err != nil {
			fmt.Println("unpack occur err")
			continue
		}

		//清空buff，否则下次读取到的数据将直接追加在末尾
		buff = make([]byte, 2048)
	}
}

func readBuff(readCh chan []byte, sendCh chan *CmdInfo) {
	for buff := range readCh {
		HandleData(buff, sendCh)
	}
}

func sendBuff(sendCh chan *CmdInfo, stopCh chan int, conn net.Conn) {
	for {
		select {
		case <-stopCh:
			return
		default:
		}

		select {
		case <-stopCh:
			return
		case cmdInfo := <-sendCh:
			pack := Package{
				CmdType: DATA,
				CmdInfo: *cmdInfo,
				Time:    GetCurrTime(),
			}
			b, err := json.Marshal(pack)
			if err != nil {
				log.Log("marshal pack failure")
				continue
			}
			conn.Write(Pack(b))
		}
	}
}

func checkErr(err error) {
	if err != nil {
		os.Exit(1)
	}
}
