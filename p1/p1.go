package p1

import (
	"fmt"

	"github.com/jogendra/ic/p2"
)

type P1 struct{}

func (p P1) HelloFromP1() {
	fmt.Println("Hello from package p1")
}

func CallingHelloFromP2() {
	p2.HelloFromP2()
}
