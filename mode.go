package main

import (
	"fmt"
	"strings"
)

//go:generate stringer -type=Mode
type Mode int

const (
	Read Mode = 1 << iota
	Write
	Execute
	FinalMode // terminator
)

func main() {
	modes := []Mode{
		Read,
		Write,
		Read | Write,
		Read | Execute,
		Read | Write | Execute,
	}

	for _, mode := range modes {
		var bs []string
		for _, b := range bits(mode) {
			bs = append(bs, b.String())
		}
		fmt.Println(mode, "=", strings.Join(bs, ", "))
	}
}

func bits(mode Mode) (modes []Mode) {
	for m := Read; m < FinalMode; m <<= 1 {
		if mode&m == m {
			modes = append(modes, m)
		}
	}
	return
}
