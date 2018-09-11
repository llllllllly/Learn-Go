package main

import "fmt"

type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	connect()
}

type PersonalCompute struct {
	name string
}

type Android struct {
	name string
}

func (pc PersonalCompute) Name() string {
	return pc.name
}

func (pc PersonalCompute) connect() {
	fmt.Println("Connect is:", pc.name)
}

func (a Android) Name() string {
	return a.name
}

func (a Android) connect() {
	fmt.Println("Connect is:", a.name)
}

func main() {
	pc := PersonalCompute{"pc"}
	Disconnect(pc)
	a := Android{"Android"}
	Disconnect(a)
}

func Disconnect(usb USB) {
	if pc, ok := usb.(PersonalCompute); ok {
		fmt.Println(pc.name)
		return
	}
	fmt.Println("Unknown device")
}