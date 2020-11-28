package main

import (
	"fmt"
	"path"
	"runtime"
)

type Uint32Container struct {
	s []uint32
}

func NewUint32Container() *Uint32Container {
	return &Uint32Container{s: []uint32{}}
}

func (c *Uint32Container) Put(val uint32) {
	c.s = append(c.s, val)
}

func (c *Uint32Container) Get() uint32 {
	r := c.s[0]
	c.s = c.s[1:]
	return r
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file) // Base函数返回路径的最后一个元素
	return
}

func main() {

}
