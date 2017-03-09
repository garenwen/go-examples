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
	ReflectTest()

}
