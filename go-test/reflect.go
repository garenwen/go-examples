package main

import (
	"encoding/json"
	"log"
)

func ReflectTest() {

	//var i int = 1
	//
	//// 获取Value,这里注意,如果你要改变这个变量的话,需要传递变量的地址
	//value := reflect.ValueOf(&i)
	//
	//// value是一个指针,这里获取了该指针指向的值,相当于value.Elem()
	//value = reflect.Indirect(value)
	//
	//// Interface是获取该value的值,返回的是一个interface对象
	//log.Println(value.Interface()) // 输出:1
	//
	//// 把变量i的值设为2
	//if value.Kind() == reflect.Int {
	//	value.SetInt(2)
	//}
	//
	//log.Println(value.Interface()) // 输出:2

	//type S struct {
	//	A string // 注意:只有大写开头的成员变量可以Set
	//}
	//
	//s := S{"x"}
	//
	//value := reflect.ValueOf(&s)
	//
	//value = reflect.Indirect(value)
	//
	//
	////value是结构体s,所以打印出来的是整个结构体的信息
	//log.Println(value.Interface()) //输出: {x}
	//
	//f0 := value.FieldByName("A") //获取结构体s中第一个元素a
	//
	//log.Println(f0) // 输出: x
	//
	//if f0.Kind() == reflect.String {
	//	if f0.CanSet() {
	//		f0.SetString("y")
	//	}
	//}
	//
	//log.Println(f0) // 输出: y
	//
	//log.Println(value.Interface()) //输出: {y}

	type S struct {
		A string
	}

	s := S{}

	r, _ := json.Marshal(s)

	log.Println(string(r))

	//log.Printf(s.A)
	//
	//value := reflect.ValueOf(&s)
	//
	//value = reflect.Indirect(value)
	//
	////获取结构体s的类型S
	//vt := value.Type()
	//
	////获取S中的A成员变量
	//f, _ := vt.FieldByName("A")

	//获取成员变量A的db标签
	//log.Println(f.Tag.Get("json")) //输出: tag_a
}
