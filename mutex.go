package ztool

import (
	"sync"
)

var mu sync.Mutex

type DB struct {
	Name string
}

func GetDb() *DB {
	mu.Lock()
	defer mu.Unlock()
	db := &DB{Name: "222"}
	return db
}

func (db *DB) Get() {
	//time.Sleep(10 * time.Millisecond)
	//log.Println()
}
