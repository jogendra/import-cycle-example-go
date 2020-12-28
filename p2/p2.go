package p2

import (
	"fmt"
	"import-cycle-example/p1"
)

type PP2 struct{}

func New() *PP2 {
	return &PP2{}
}

func (p *PP2) HelloFromP2() {
	fmt.Println("Hello from package p2")
}

func (p *PP2) HelloFromP1Side() {
	pp1 := p1.New()
	pp1.HelloFromP1()
}
