package main

import (
	"graphs/graphs"
	"fmt"
)

func main() {
	g := [][]int{
		{ 0, 2, 0, 6, 0 }, 
		{ 2, 0, 3, 8, 5 }, 
		{ 0, 3, 0, 0, 7 }, 
		{ 6, 8, 0, 0, 9 }, 
		{ 0, 5, 7, 9, 0 },
	}

	err := graphs.PrimMST(g)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
	}
}