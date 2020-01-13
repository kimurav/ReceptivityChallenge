package test

import (
	"testing"

	"github.com/manavp1996/receptivity/data"
)

func TestMatrix(t *testing.T) {
	matrix, err := data.NewMatrix()
	if err != nil {
		t.Errorf("could not create vertex matrix: %v", err)
	}
	t.Run("Q1", func(t *testing.T) {
		weight := matrix.GetRouteWeight([]int{0, 1, 2})
		if weight != 9 {
			t.Errorf("Q1: expected 9 got %d", weight)
		}
	})
	t.Run("Q2", func(t *testing.T) {
		weight := matrix.GetRouteWeight([]int{0, 3})
		if weight != 5 {
			t.Errorf("Q2: expected 5 got %d", weight)
		}
	})
	t.Run("Q3", func(t *testing.T) {
		weight := matrix.GetRouteWeight([]int{0, 3, 2})
		if weight != 13 {
			t.Errorf("Q3: expected 13 got %d", weight)
		}
	})
	t.Run("Q4", func(t *testing.T) {
		weight := matrix.GetRouteWeight([]int{0, 4, 1, 2, 3})
		if weight != 22 {
			t.Errorf("Q4: expected 14 got %d", weight)
		}
	})
	t.Run("Q5", func(t *testing.T) {
		weight := matrix.GetRouteWeight([]int{0, 4, 3})
		if weight != 0 {
			t.Errorf("Q5: expected 0 got %d", weight)
		}
	})
	t.Run("Q6", func(t *testing.T) {
		matrix.CleanStack()
		matrix.GetCycle(2, 2)
		if matrix.Q6 != 2 {
			t.Errorf("Q6: expected 2 got %d", matrix.Q6)
		}
	})
	t.Run("Q7", func(t *testing.T) {
		matrix.CleanStack()
		matrix.FindPath(0, 2, 0)
		if matrix.Q7 != 3 {
			t.Errorf("Q7: expected 3 got %d", matrix.Q7)
		}
	})
	t.Run("Q8", func(t *testing.T) {
		paths := make([]int, 0, 10)
		matrix.ShortestPath(0, 2, 0, &paths)
		min, err := data.FindMin(paths)
		if err != nil {
			t.Errorf("failed to find minimum path: %v", err)
		}
		if min != 9 {
			t.Errorf("Q8: expected 9 got %d", min)
		}
	})
}
