package storage

import (
	"os"
	"strconv"
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

func SafeAdd(key string) {
	mux.Lock()
	var file *os.File
	var content []byte

	defer mux.Unlock()
	defer file.Close()
	file = getWithCreateIfNotExists(numFileName)
	file.Read(content)
	if con
	num, _ := strconv.Atoi(string(content[:]))
	num++
}

//type Storage struct {
//	source io.ReadWriter
//	mutex  *sync.RWMutex
//}
//
//func (s *Storage) Add(key string) {
//	s:= os.Open()
//}
