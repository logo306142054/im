/*
* created by Orz on 2017/6/9
 */
package public

import (
	"bytes"
	"encoding/binary"
	"errors"
	"reflect"
	"strconv"
)

//数据包的类型
const (
	INVALID_TYPE = iota
	CONTROL      //控制包
	DATA         //数据包
)

type CmdInfo struct {
	Cmd     int16  `json:"cmd"`
	Err     int    `json:"errcode`
	ErrInfo string `json:"errinfo"`
	Data    []byte `json:"data", omitempty`
}

type Package struct {
	CmdType byte    `json:"ct"`
	CmdInfo CmdInfo `json:"ci"`
	Time    string  `json:"time"`
}

var (
	//包头内容
	HEAD = "myfirstIMproject."
	//包头的长度
	HEAD_BYTE_LEN int = len(HEAD)
	//包版本
	VER_NUM_BYTE_LEN int = int(reflect.TypeOf(int16(0)).Size())
	//实际数据内容的长度
	DATA_LEN_BYTE_LEN int = int(reflect.TypeOf(int16(0)).Size())
	//包的最小固定长度
	FIX_HEAD_LEN = HEAD_BYTE_LEN + VER_NUM_BYTE_LEN + DATA_LEN_BYTE_LEN
)

var (
	CURRENT_HEAD_VER = 1
)

var emptyBytes = make([]byte, 0)

func Pack(buff []byte) []byte {
	newBuff := make([]byte, 0)
	newBuff = append(newBuff, []byte(HEAD)...)
	newBuff = append(newBuff, ParseInt16ToBytes(CURRENT_HEAD_VER)...)
	newBuff = append(newBuff, ParseInt16ToBytes(len(buff))...)
	newBuff = append(newBuff, buff...)
	return newBuff
}

func Unpack(buff []byte, readCh chan []byte) ([]byte, error) {
	totalLen := len(buff)
	pos := 0
	for (pos + FIX_HEAD_LEN) < totalLen {
		//校验包头
		if !isHead(buff[pos : pos+HEAD_BYTE_LEN]) {
			return emptyBytes, errors.New("invalid package head")
		}

		pos += HEAD_BYTE_LEN

		//检查包版本
		verNum := getVerNum(buff[pos : pos+VER_NUM_BYTE_LEN])
		if verNum != CURRENT_HEAD_VER {
			return emptyBytes, errors.New("invalid package version")
		}

		pos += VER_NUM_BYTE_LEN

		//获取包真实数据的长度
		dataLen := getDataLen(buff[pos : pos+DATA_LEN_BYTE_LEN])
		pos += DATA_LEN_BYTE_LEN

		//数据长度有误，丢弃
		if pos+dataLen > totalLen {
			return emptyBytes, errors.New("data length is too long, length=" +
				strconv.Itoa(pos+dataLen) + " totallen=" + strconv.Itoa(totalLen))
		}

		data := buff[pos : pos+dataLen]
		readCh <- data

		pos += dataLen
	}
	//pos != countLen说明还有数据没读完，但剩余的数据无法组成完整的包
	//需保存这部分数据，并拼接到下一次读取到的数据的前部
	if pos != totalLen {
		return buff[pos:], nil
	}
	//数据解析完，返回一个空的数组
	return emptyBytes, nil
}

func isHead(buff []byte) bool {
	if string(buff) == HEAD {
		return true
	}
	return false
}

func getVerNum(buff []byte) int {
	return GetInt16FromBuff(buff)
}

func getDataLen(buff []byte) int {
	return GetInt16FromBuff(buff)
}

func GetInt16FromBuff(buff []byte) int {
	var data int16
	byteBuffer := bytes.NewBuffer(buff)
	binary.Read(byteBuffer, binary.BigEndian, &data)

	return int(data)
}

func ParseInt16ToBytes(digit int) []byte {
	tmp := int16(digit)
	buff := make([]byte, 0)
	byteBuff := bytes.NewBuffer(buff)
	binary.Write(byteBuff, binary.BigEndian, tmp)

	return byteBuff.Bytes()
}
