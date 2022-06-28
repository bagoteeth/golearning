package main

import (
	"fmt"
	"strings"
)

func main() {
	b := []string{"aaa", "bbb", "ccc"}
	str := `[
    {
        "$match":{
            "$and":[
                {"$expr": {"$in": ["$metadata.namespace", %s]}},
                {"$expr": {"$in": ["$spec.clusterName", ["rfjiang111","aaaa"]]}}
            ]
        }
    },{
        "$project":{
            _id:0,
            Name: "$metadata.Name",
            appType: "$kind",
            clusterName: "$spec.clusterName",
            namespace: "$metadata.namespace",
            uids: {
                $map:{
                    input: "$subResourceReferences",
                    as : "uids",
                    in:"$$uids.uid"
                }
            }
        }
    }
]`
	c := strings.Join(b, "\",\"")
	c = "[\"" + c + "\"]"

	fmt.Printf(str, c)
}
