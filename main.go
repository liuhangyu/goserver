package main

import (
	"fmt"
	"sync"
	//. "goserver/channel"
	. "goserver/operfile"
)

type ElemValue struct {
	data string
}

type DbaseSet struct {
	imap map[string]*ElemValue
	//lock sync.Mutex
	lock sync.RWMutex
}

func (db *DbaseSet) Init(size uint64) {
	db.imap = make(map[string]*ElemValue)
	for i := uint64(0); i < size; i++ {
		key := fmt.Sprintf("%#x", i)
		value := new(ElemValue)
		value.data = fmt.Sprintf("%#x", i)
		db.imap[key] = value
	}
}

func (db *DbaseSet) DisplayDb() {
	fmt.Printf("Show map data\n")
	for k, v := range db.imap {
		fmt.Printf("key:%s\tvalue:%s\n", k, v.data)
	}
}

func (db *DbaseSet) Write(key *string, value *ElemValue) {
	if key == nil || value == nil {
		fmt.Printf("%s\n", "input parmter abort!")
	}
	//db.lock.Lock() // lock the mutex
	//defer db.lock.Unlock() // unlock the mutex at the end
	db.lock.Lock()
	defer db.lock.Unlock()

	db.imap[*key] = value
}

func (db *DbaseSet) Read(key string) *ElemValue {
	//db.lock.Lock() // lock the mutex
	//defer db.lock.Unlock() // unlock the mutex at the end
	db.lock.RLock()
	defer db.lock.RUnlock()

	for k, v := range db.imap {
		if k == key {
			return v
		}
	}
	return nil
}

func (db *DbaseSet) FmtWrite(k uint64, v uint64) {
	key := fmt.Sprintf("%#x", k)
	value := new(ElemValue)
	value.data = fmt.Sprintf("%#x", v)
	db.Write(&key, value)
}

func (db *DbaseSet) FmtRead(k uint64) string {
	key := fmt.Sprintf("%x", k)
	elem := db.Read(key)
	if elem != nil {
		return elem.data
	}
	return ""
}

const (
	size = 1000000
)

/*
func main() {
	db := new(DbaseSet)
	db.Init(size)
	db.DisplayDb()
	value := db.FmtRead(100)
	fmt.Printf("value:%s\n", value)

	for i := uint64(0); i < size+1000; i++ {
		go db.FmtWrite(i, i)
	}

	for i := uint64(0); i < 1000; i++ {
		go db.FmtRead(i)
	}
}



func main(){
	ChannelFunc()
}

*/
func main(){
	Test()
	
}