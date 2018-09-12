package main

import (
	"fmt"
	"reflect"
)

type Log struct {
	username string
	password string
}

type User struct {
	Log
	Name 	string
	Age		int
}

func (u User) LogIn(username string, password string) bool {
	return u.username == username && u.password == password
}

func (u *User) LogUp(username string, password string, Name string, Age int) {
	u.username 	= username
	u.password 	= password
	u.Name 		= Name
	u.Age 		= Age
}

func main() {
	u := User{}
	LogUp(&u)
	ShowInfo(u, LogIn(u))
}

func LogUp(o interface{}) {
	v := reflect.ValueOf(o)
	m := v.MethodByName("LogUp")
	if v.Kind() != reflect.Ptr || !m.IsValid(){
		fmt.Println("注册失败")
		return 
	}
	args := []reflect.Value{
		reflect.ValueOf("llllllllly"),
		reflect.ValueOf("123456"),
		reflect.ValueOf("l"),
		reflect.ValueOf(18),
	 }
	m.Call(args)
	fmt.Println("注册成功")
}

func LogIn(o interface{}) bool{
	v := reflect.ValueOf(o)
	m := v.MethodByName("LogIn")
	if v.Kind() != reflect.Struct && !m.IsValid() {
		fmt.Println("登录失败")
		return false
	}
	args := []reflect.Value{
		reflect.ValueOf("llllllllly"),
		reflect.ValueOf("123456"),
	}
	result := m.Call(args)[0]
	return result.Bool()
}

func ShowInfo(o interface{}, isLogIn bool) {
	if !isLogIn {
		fmt.Println("登录失败")
		return
	}
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)

	for i := 0; i < t.NumField(); i++ {
		f 	:= t.Field(i)
		if f.Anonymous {
			continue
		}
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
}