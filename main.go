package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/tameralamiri/go-hello-world/morestrings"
)

func main() {
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
