package main

import (
	"bytes"
	"sync"
)

// map 非线程安全

type Info struct {
	mu sync.Mutex
	a  string
}

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func Update(info *Info) {
	info.mu.Lock()
	info.a = "aaa"
	info.mu.Unlock()
}

func Update2(syncBuffer *SyncedBuffer) {
	syncBuffer.lock.Lock()

}
