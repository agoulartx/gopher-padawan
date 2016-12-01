package main

import (
	"fmt"
	"github.com/agoulartx/gopher-padawan/scripts/04_scope_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
