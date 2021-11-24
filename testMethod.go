package main

import "fmt"

type Receiver struct {
	name string
}

func (r *Receiver) dosth() {
	fmt.Println("i am " + r.name)
}
func (r *Receiver) setName(name string) {
	r.name = name
}
func (r *Receiver) setName2(name string) {
	r.name = name
}

func main() {
	r := Receiver{name: "receiver"}
	r.dosth()
	r.setName("pointreceiver")
	r.dosth()
	r.setName2("valuereceiver")
	r.dosth()
}
