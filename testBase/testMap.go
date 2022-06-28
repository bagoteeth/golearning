package main

import "fmt"

type testmapstr struct {
	Name string
	Age  int
}

func (t *testmapstr) dosth() {
	fmt.Printf("i'm %s and i'm %d\n", t.Name, t.Age)
	t.Age = 88
}

var mmp map[string]*testmapstr

func main() {
	mmp = make(map[string]*testmapstr)
	a, ok := mmp["aaa"]
	if !ok {
		a = &testmapstr{Name: "aaaaa"}
		mmp["aaa"] = a
	}
	fmt.Println(a)
	fmt.Println(mmp["aaa"])
	a.dosth()
	fmt.Println(a)
	fmt.Println(mmp["aaa"])

}
