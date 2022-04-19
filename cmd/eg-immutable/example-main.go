package main

import (
	"fmt"

	"github.com/fliedonion/eg-go-immutable/internal/immutable"
)

func main() {

	x := immutable.NewMyImmutable(1)
	// x.val = 0 // error.
	fmt.Println(x.Value())

	y := immutable.NewMyImmutable("string value")
	fmt.Println(y.Value())
}
