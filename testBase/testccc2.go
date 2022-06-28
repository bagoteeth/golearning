package main

import (
	"fmt"
	"sort"
)

const (
	Cluster   = "cluster"
	Namespace = "namespace"
	App       = "app"
)

func main() {
	fmt.Println(Cluster > Namespace)
	a := []string{App, Namespace, App, Cluster, Namespace}

	fmt.Println(a)
	sort.SliceStable(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	fmt.Println(a)
}

//func judge(a,b string) bool{
//
//}
