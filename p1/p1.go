package p1

import (
	"fmt"
	_ "unsafe"
)

type PP1 struct{}

func (p *PP1) HelloFromP1() {
	fmt.Println("Hello from package p1")
}

//go:noescape
//go:linkname helloFromP2A p2.helloFromP2A
func helloFromP2A()

func (p *PP1) HelloFromP2Side() {
	helloFromP2A()
}
