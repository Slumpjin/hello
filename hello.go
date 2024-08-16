package main

import (
	"fmt"

	"github.com/Slumpjin/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG, olleH"))
	fmt.Println(cmp.Diff("Hello, world", "hello, go"))
}
