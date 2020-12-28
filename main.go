package main

import (
	"import-cycle-example/p1"
)

func main() {
	pp1 := p1.PP1{}
	pp1.HelloFromP2Side()
}
