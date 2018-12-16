package lib

import "fmt"

type printable interface {
	String() string
}

// Print ...
func Print(x printable) {
	fmt.Printf("// %s\n", x.String())
}
