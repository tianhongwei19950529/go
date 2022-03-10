package main

import (
	"fmt"
	"sync"
	"unsafe"
)

type Students struct {
	name string
}

var singleInstance *Students
var once sync.Once

func newStudents() *Students {
	once.Do(func() {
		singleInstance = new(Students)
	})
	return singleInstance
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(){
			obj := newStudents()
			fmt.Printf("%X\n",unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
