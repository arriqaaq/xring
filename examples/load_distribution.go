package main

/*
	Examples taken from "github.com/buraksezer/consistent"
*/

import (
	"fmt"
	"github.com/arriqaaq/xring"
	"math/rand"
)

func main() {
	nodes := []string{"a", "b", "c"}
	cnf := &xring.Config{
		VirtualNodes: 100,
		LoadFactor:   2,
	}
	hashRing := xring.NewRing(nodes, cnf)

	keyCount := 1000000
	distribution := make(map[string]int)
	key := make([]byte, 4)
	for i := 0; i < keyCount; i++ {
		rand.Read(key)
		node, err := hashRing.Get(string(key))
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		hashRing.Done(node)
		distribution[node]++
	}
	for node, count := range distribution {
		fmt.Printf("node: %s, key count: %d\n", node, count)
	}
}
