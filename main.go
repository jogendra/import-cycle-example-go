package main

import (
	"import-cycle-example/p1"
	"import-cycle-example/p2"
)

func main() {
	pp1 := p1.PP1{}
	pp2 := p2.New(&pp1)
	pp2.HelloFromP1Side()
}
