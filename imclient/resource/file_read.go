package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"syscall"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestReadFile() {

	//	file, err := os.Create("dir.txt")
	//	check(err)
	//	defer file.Close()

	//	w := bufio.NewWriter(file)
	//	w.WriteString("haha")
	//	w.Flush()

	buf, err := ioutil.ReadFile("dir.txt")
	check(err)
	fmt.Println(string(buf))

	fileinfo, err := ioutil.ReadDir(".")
	for _, v := range fileinfo {
		fmt.Println(v.Name(), v.ModTime().Format("2006-01-02 15:04:05"), v.Size(), v.Mode())
	}

	ioutil.WriteFile("dir.txt", []byte("laddla"), 0644)

	file, err := os.OpenFile("dir.txt", os.O_APPEND|os.O_RDWR, 0644)
	check(err)
	defer file.Close()
	file.WriteString("append content")
	file.Sync()

	file.Seek(0, syscall.FILE_BEGIN)
	br := bufio.NewReader(file)

	line, isPrefix, err := br.ReadLine()
	check(err)
	if isPrefix {
		fmt.Println("prefix")
		return
	}
	fmt.Println(string(line))

	s := []int{5, 4, 3, 2, 1}
	//qs := []int{40, 4, 30, 50, 35, 84, 32, 18, 40}
	qs := []int{40, 50, 60}
	bubbleSort(s)
	quickSort(qs, 0, len(qs)-1)
	fmt.Println(qs)
}

func bubbleSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func quickSort(s []int, left int, right int) {
	base := s[left]
	L := left
	R := right

	for L < R {
		for R > L && s[R] >= base {
			R--
		}
		if R > L {
			s[L] = s[R]
			L++
		}

		for L < R && s[L] <= base {
			L++
		}
		if L < R {
			s[R] = s[L]
			R--
		}
	}
	s[L] = base

	if left < R-1 {
		quickSort(s, left, R-1)
	}
	if L+1 < right {
		quickSort(s, L+1, right)
	}
}

type MyInteface interface {
	Read()
}
type Base struct {
	v    int
	read func()
}

func r() {
	fmt.Println("r test")
}

func (b *Base) Read() {
	fmt.Println("Read interface")
}

type MyError struct {
	ErrMsg string
	Err    error
}

func (me *MyError) Error() string {
	return me.ErrMsg + me.Err.Error() + " " + reflect.TypeOf(me.Err).Elem().String()
}

func NewMyError(msg string, err error) error {
	me := &MyError{
		ErrMsg: msg,
		Err:    err,
	}
	return me
}

func main() {
	resp, _ := http.Head("http://www.baidu.com")
	//io.Copy(os.Stdout, resp.Header)
	defer resp.Body.Close()
	fmt.Println(resp.Header)

	fmt.Println(resp.Header.Get("Connection"))

	b := Base{read: r}
	b.read()
	var myin MyInteface = &b
	myin.Read()

	err := errors.New("lala")
	me := NewMyError("my error ", err)
	log.Fatal(me)

	TestReadFile()

}
