package main

import (
	"fmt"

	data "github.com/manavp1996/receptivity/data"
)

func main() {
	matrix, err := data.NewMatrix()
	if err != nil {
		fmt.Printf("Failed to build matrix")
	}
	//matrix.PrintGraph()
	arr := []int{0, 1, 2}
	fmt.Printf("Q1 path A-B-C weight: %d \n", matrix.GetRouteWeight(arr))
	fmt.Printf("Q2 path A-D weight is %d\n", matrix.GetRouteWeight([]int{0, 3}))
	fmt.Printf("Q3 path A-D-C weight is %d\n", matrix.GetRouteWeight([]int{0, 3, 2}))
	fmt.Printf("Q4 path A-E-B-C-D weight is %d\n", matrix.GetRouteWeight([]int{0, 4, 1, 2, 3}))
	fmt.Printf("Q5 path A-E-D weight is %d\n", matrix.GetRouteWeight([]int{0, 4, 3}))
	matrix.Stack.PushFront(0)
	matrix.GetCycle(2, 2)
	matrix.CleanStack()
	matrix.FindPath(0, 2, 0)
	fmt.Printf("Q7 the number of paths are: %d\n", matrix.Q7)

	paths := make([]int, 0, 10)
	matrix.ShortestPath(0, 2, 0, &paths)
	if min, err := data.FindMin(paths); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Q8: %d\n", min)
	}

}
