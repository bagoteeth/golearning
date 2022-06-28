package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	var a = `
{
    "clusters": [
        {
            "Name": "cluster1"
        },
        {
            "Name": "cluster2"
        }
    ],
    "hostGroups": [
        {
            "Name": "default",
            "clusterName": "cluster1"
        }
    ],
    "hosts": [
        {
            "Name": "host1",
            "hostGroupName": "default",
            "clusterName": "cluster1"
        }
    ],
    "containers": [
        {
            "containerID": "",
            "hostName": "",
            "hostGroupName": "",
            "clusterName": ""
        }
    ],
    "namespaces": [
        {
            "Name": "",
            "clusterName": ""
        }
    ],
    "apps": [
        {
            "Name": "",
            "kind": "",
            "namespace": "",
			"clusterName": "fffffff"
        }
    ],
    "services": [
        {
            "Name": "",
            "hostName": "",
            "hostGroupName": "",
            "namespace": "",
            "clusterName": ""
        }
    ]
}`

	var rnm map[string][]map[string]string
	//var b interface{}
	json.Unmarshal([]byte(a), &rnm)
	fmt.Println(reflect.TypeOf(rnm))
	fmt.Println(rnm)
	//fmt.Println(reflect.TypeOf(b))
	//x,ok:=b.(map[string]interface{})
	//if !ok{
	//	fmt.Println("nononono1")
	//	return
	//}
	//fmt.Println(reflect.TypeOf(x))
	//
	//y,ok:=x["apps"].([]interface{})
	//if !ok{
	//	fmt.Println("nononono2")
	//	return
	//}
	//fmt.Println(reflect.TypeOf(y))
	//
	//z,ok:=y[0].(map[string]interface{})
	//if !ok{
	//	fmt.Println("nononono3")
	//	return
	//}
	//fmt.Println(reflect.TypeOf(z))
	//
	//h,ok:=z["clusterName"].(string)
	//if !ok{
	//	fmt.Println("nononono4")
	//	return
	//}
	//fmt.Println(reflect.TypeOf(h))
	//fmt.Println(h)
}
