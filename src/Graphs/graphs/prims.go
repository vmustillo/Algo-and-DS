package graphs

import (
	"math"
	"errors"
	"fmt"
)

func GetMinimumEdge(visited []bool, edges []int) int {
	min := math.MaxInt64
	var indexMinEdge int
	
	for i := 0; i < len(visited); i++ {
		if visited[i] == false && edges[i] < min {
			min = edges[i]
			indexMinEdge = i
		}
	}

	return indexMinEdge
}

func PrintMST(parents []int, graph [][]int) {
	fmt.Println("Vertex\t\tLength")

	for i := 1; i < len(parents); i++ {
		fmt.Printf("%d - %d\t\t  %d\n", parents[i], i, graph[i][parents[i]])
	}
}

func PrimMST(graph [][]int) error {
	graphRows := len(graph)

	if graphRows == 0 {
		err := errors.New("Cannot find PrimMST: empty graph provided")

		return err
	}

	edges := make([]int, graphRows)
	parents := make([]int, graphRows)
	visited := make([]bool, graphRows) // all false by default
	
	// initialize edges to maximum value, 
	for i := 1; i < graphRows; i++ {
		edges[i] = math.MaxInt64;
	}

	// initialize first vertex to have no parent and 0 edge length
	edges[0] = 0
	parents[0] = -1

	for i := 0; i < graphRows; i++ {
		row := GetMinimumEdge(visited, edges)
		
		visited[row] = true

		for col := 0; col < graphRows; col++ {
			if graph[row][col] != 0 && visited[col] == false && graph[row][col] < edges[col] {
				fmt.Printf("graph[%d][%d] = %d\n", row, col, graph[row][col])
				edges[col] = graph[row][col]
				parents[col] = row
			}
		}
	}

	PrintMST(parents, graph)

	return nil
}