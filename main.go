package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	_ "net/http/pprof"
	"runtime"
)

func worker(ch chan struct{}) {
	<-ch
	println("roger1")
	// send a message to the main program
	close(ch)
}

var intMap map[int]int
var cnt = 8192

func initMap() {
	intMap = make(map[int]int, cnt)

	for i := 0; i < cnt; i++ {
		intMap[i] = i
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v, totalAlloc=%v, sys=%v, numGc=%v", m.Alloc, m.TotalAlloc, m.Sys, m.NumGC)
}

type Q struct {
	Name string
}

func (q *Q) sayHello() bool {
	fmt.Printf("%p\n", q)
	return q == nil
}

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)
	return sData
}

type DeviceInfo struct{}

func main() {
	ginEngin := gin.New()
	ginEngin.GET("/test", func(ctx *gin.Context) {
		var d *DeviceInfo
		ctx.Set("__di", d)
		info, exist := ctx.Get("__di")
		fmt.Println(info, exist)
		ctx.JSON(200, map[string]interface{}{"info": info, "exist": exist})
	})
	ginEngin.Run("9000")
}
