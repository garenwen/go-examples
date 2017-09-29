package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func f1() {
	var i = 1
	for {
		i++
	}
}
func f2() {
	for {
		fmt.Println("f2")
		time.Sleep(time.Second)
	}
}

func gr() {

	log.Println("open cpu profile.")
	f, err := os.Create("mallcenter_cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}

	defer pprof.StopCPUProfile()

	go f1()
	go f1()
	go f1()
	go f1()
	go f2()
	go f2()
	go f2()
	go f2()

	for {
		fmt.Println("main")
		time.Sleep(time.Second)
	}
}
func main() {
	runtime.GOMAXPROCS(4)
	//QRCode()
	//Merge()
	//DownLoad()
	//gr()
	//ReflectTest()

	//新建一个logger
	//logger.New("mallcenter")
	//
	////打印Info信息
	//logger.Info("test_topic", map[string]interface{}{"key": 2})
	//
	////打印警告信息
	//logger.Warnning("warnning log")
	//
	////打印debug信息
	//logger.Debug("debug log")
	//
	////打印error信息
	//logger.Error("error log")

	//QRCode()
	//Merge()

	etcdmain()
}
