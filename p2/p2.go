package p2

import (
	"fmt"

	"github.com/jogendra/ic/p1"
)

type P2 struct{}

func (p P2) HelloFromP2() {
	fmt.Println("Hello from package p2")
}

func CallingHelloFromP1() {
	p1.CallingHelloFromP2()
}
