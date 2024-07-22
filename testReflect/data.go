package main

import "strings"

type FilterIntf interface {
	A() bool
	B() bool
}

type FilterA struct {
	Name string
}

type FilterB struct {
	Name string
}

func (r *FilterA) A() bool {
	return strings.Contains(r.Name, "a")
}

func (r *FilterA) B() bool {
	return false
}

func (r *FilterB) A() bool {
	return strings.Contains(r.Name, "b")
}

func (r *FilterB) B() bool {
	return false
}

type BigFilter struct {
	AAA string
	BBB string
	FA  []FilterA
	FB  []FilterB
}
