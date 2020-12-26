package p1

import (
	"fmt"

	"import-cycle-example/p2"
)

func HelloFromP1() {
	fmt.Println("Hello from package p1")
}

func HelloFromP2Side() {
	p2.HelloFromP2()
}
