package p2

import (
	"fmt"

	"import-cycle-example/p1"
)

func HelloFromP2() {
	fmt.Println("Hello from package p2")
}

func HelloFromP1Side() {
	p1.HelloFromP1()
}
