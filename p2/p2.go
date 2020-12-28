package p2

import (
	"fmt"
)

type pp1 interface {
	HelloFromP1()
}

type PP2 struct {
	PP1 pp1
}

func New(pp1 pp1) *PP2 {
	return &PP2{
		PP1: pp1,
	}
}

func (p *PP2) HelloFromP2() {
	fmt.Println("Hello from package p2")
}

func (p *PP2) HelloFromP1Side() {
	p.PP1.HelloFromP1()
}
