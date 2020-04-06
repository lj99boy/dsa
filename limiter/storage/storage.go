package storage

import (
	"bytes"
	"encoding/binary"
	"os"
	"sync"
)

var storageFileName string
var numFileName string

func init() {
	storageFileName = "./default.bak"
	numFileName = "./numDefault.bak"
}

var mux *sync.RWMutex

func SetFileName(name string) {
	storageFileName = name
}

func getWithCreateIfNotExists(fileName string) *os.File {
	_, err := os.Stat(fileName)
	var file *os.File
	if os.IsNotExist(err) {
		file, _ = os.Create(fileName)
	} else {
		file, _ = os.Open(fileName)
	}
	return file
}

func SafeAdd() uint64 {
	mux.Lock()
	defer mux.Unlock()

	var content []byte
	var num uint64
	file, _ := os.OpenFile(numFileName, os.O_CREATE|os.O_RDWR, 0644)
	defer file.Close()
	file.Read(content)
	if len(content) == 0 {
		num = 0
	} else {
		num = binary.BigEndian.Uint64(content)
	}
	num++
	buf := bytes.NewBuffer(content)
	binary.Write(buf, binary.BigEndian, num)
	file.Write(content)
	return num
}

//type Storage struct {
//	source io.ReadWriter
//	mutex  *sync.RWMutex
//}
//
//func (s *Storage) Add(key string) {
//	s:= os.Open()
//}
