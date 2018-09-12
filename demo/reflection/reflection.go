package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id 		int
	Name 	string
	Age		int
}

type Manager struct {
	User
	title string
}

func (u User) Hello(name string) {
	fmt.Println("Hello", name, ", my name is:", u.Name)
}

func main() {
	//setUser()
	setManager()
}

func setUser() {
	u := User{1, "l", 10}
	Info(u)
	fmt.Println(u)
	Set(&u)
	fmt.Println(u)
}

func setManager() {
	m := Manager{User: User{1, "l", 10}, title: "ok"}
	m.Hello("hhh")
	//Anonymous(m)
	call(m)
}

func call(o interface{}) {
	v := reflect.ValueOf(o)
	m := v.MethodByName("Hello1")
	if v.Kind() != reflect.Struct || !m.IsValid() {
		fmt.Println("xxx")
		return
	}
	args := []reflect.Value{reflect.ValueOf("hhh")}
	m.Call(args)
}

func Anonymous(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() != reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")

	if !f.IsValid() {
		fmt.Println("xxx")
		return 
	}
	if f.Kind() == reflect.String {
		f.SetString("OK")
	}

}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	fmt.Println("Methods:")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}