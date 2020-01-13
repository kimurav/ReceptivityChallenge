package data

import (
	"container/list"
	"errors"
	"fmt"
)

const matrix_size = 5

type Graph struct {
	Matrix  [5][5]int
	Stack   *list.List
	Visited []Node
	Q7      int
	Q8      int
	Q6      int
}

/*Add a pair of verticies with an edge weighting to the matrix*/
func (m *Graph) AddNode(x, y, w int) error {
	if x > matrix_size {
		return errors.New("value of x is greater than 5")
	}
	if y > matrix_size {
		return errors.New("value of y is greater than 5")
	}
	m.Matrix[x][y] = w
	return nil
}

/*Add a Node to the visited list*/
func (g *Graph) AddVisitedNode(source, dest int) {
	n := Node{
		Source: source,
		Dest:   dest,
	}
	g.Visited = append(g.Visited, n)
}

/*Utility function to print the matrix*/
func (g Graph) PrintGraph() {
	for i := 0; i < matrix_size; i++ {
		for j := 0; j < matrix_size; j++ {
			fmt.Printf("%d ", g.Matrix[i][j])
		}
		fmt.Println()
	}
}

/*Q1-5 getting the total weights for the routes*/
func (g Graph) GetRouteWeight(route []int) int {
	numStops := len(route)
	weight := 0
	i := 0
	for j := 1; j < numStops; j++ {
		if g.Matrix[route[i]][route[j]] == 0 {
			return 0
		}
		weight += g.Matrix[route[i]][route[j]]
		i++
	}
	return weight
}

/*Finding the shortest path between two vertices. Makes use of a
stack and is used for Q7*/
func (g *Graph) FindPath(startPoint, endPoint, currentStop int) {
	if g.Stack.Len() > 4 {
		return
	}
	if n := g.Stack.Front(); n != nil {
		if currentStop == endPoint && g.Stack.Len() == 4 {
			g.Q7++

		}
	}
	for i := 0; i < matrix_size; i++ {
		if g.Matrix[currentStop][i] != 0 {
			g.Stack.PushFront(i)
			g.FindPath(startPoint, endPoint, i)
		}
	}
}

/*REturn the proper value or send it in as a param rempve duplicates*/
func (g *Graph) GetCycle(startPoint, currentStop int) {
	if g.Stack.Len() > 3 {
		return
	}
	if n := g.Stack.Front(); n != nil {
		if n.Value == startPoint {
			//printRoute(g.Stack)
			g.Q6 = g.Stack.Len()
			return
		}
	}
	for i := 0; i < matrix_size; i++ {
		if g.Matrix[currentStop][i] == 0 {
			continue
		}
		g.Stack.PushFront(i)
		g.GetCycle(startPoint, i)
	}
}
func (g *Graph) ShortestCycle(start, prevStop, currentStop, total int, allPaths *[]int) {
	for i := 0; i < matrix_size; i++ {
		if g.Matrix[start][i] != 0 {
			if currentStop == start {
				*allPaths = append(*allPaths, total)
				total = 0
			} else {
				if g.ContainsNode(prevStop, currentStop) {
					continue
				} else {
					g.AddVisitedNode(prevStop, currentStop)
					g.ShortestCycle(start, currentStop, i, total, allPaths)
				}
			}
		}
	}
}

func (g Graph) ContainsNode(source, dest int) bool {
	if len(g.Visited) == 0 {
		return false
	}
	for _, val := range g.Visited {
		if val.Source == source && val.Dest == dest {
			return true
		}
	}
	return false
}

/* Find the shortest path between two nodes used for Q8*/
func (g *Graph) ShortestPath(start, end, total int, allPaths *[]int) {
	for i := 0; i < matrix_size; i++ {
		if g.Matrix[start][i] != 0 {
			total += g.Matrix[start][i]
			if i == end {
				*allPaths = append(*allPaths, total)
				total = 0
			} else {
				g.ShortestPath(i, end, total, allPaths)
			}
		}
	}
}
func NewMatrix() (*Graph, error) {
	graph := Graph{}
	edgeList := CreateEdgeLists()
	for i := 0; i < len(edgeList); i++ {
		if err := graph.AddNode(edgeList[i].Source, edgeList[i].Dest, edgeList[i].Weight); err != nil {
			return nil, err
		}
	}
	graph.Stack = list.New()
	return &graph, nil
}

func printRoute(route *list.List) {
	for route.Len() > 0 {
		elem := route.Remove(route.Back())
		fmt.Printf("%d->", elem)
	}
	fmt.Println()
}

func FindMin(arr []int) (int, error) {
	if len(arr) <= 0 {
		return 0, errors.New("empty array")
	}
	min := arr[0]
	for _, val := range arr {
		if val <= min {
			min = val
		}
	}
	return min, nil
}

func (g *Graph) CleanStack() {
	for g.Stack.Len() > 0 {
		g.Stack.Remove(g.Stack.Front())
	}
}
