package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"sync"
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

type Entity struct {
}

func (a Entity) FUNC1() {
	log.Println("a")
}

type Base interface {
	FUNC1()
}

func Test(a Base) map[string]Base {
	m := make(map[string]Base)

	m["a"] = a

	return m

}

type UserConfig struct {
	SMSNotify bool
	MSGNotify bool
}

func (u *UserConfig) GetSMSNotify() bool {
	return u.SMSNotify
}

func (u *UserConfig) GetMSGNotify() bool {
	return u.MSGNotify
}

type FC func(u IUser) bool

type IUser interface {
	GetSMSNotify() bool
	GetMSGNotify() bool
}

func check(list []IUser, fc FC) []IUser {
	newList := make([]IUser, 0)
	for _, item := range list {
		if fc(item) {
			newList = append(newList, item)
		}
	}

	return newList
}

type aaa struct {
	SSS int
}

type MeetingsCache struct {
	cache map[string][]aaa
	mu    sync.RWMutex
}

var meetingsCache MeetingsCache

func InitMemCache() {
	meetingsCache = MeetingsCache{cache: map[string][]aaa{}, mu: sync.RWMutex{}}

}

func GetMeetingsByYear(year string) (ms []aaa) {
	meetingsCache.mu.RLock()
	defer meetingsCache.mu.RUnlock()
	return meetingsCache.cache[year]
}

func ClearMeetingCache() {
	meetingsCache.mu.Lock()
	defer meetingsCache.mu.Unlock()

	meetingsCache.cache = map[string][]aaa{}
	return
}

var t int

func init() {
	var err error
	t, err = strconv.Atoi("2")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("init:", t)
}

func main() {
	Merge()
}
