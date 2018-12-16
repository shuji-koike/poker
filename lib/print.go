package lib

import "fmt"

type printable interface {
	ToString() string
}

// Print ...
func Print(x printable) {
	fmt.Printf("// %s\n", x.ToString())
}
