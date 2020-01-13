package data

type Node struct {
	Source int
	Dest   int
	Weight int
}

//Graph: AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7
func CreateEdgeLists() [9]Node {
	n := [9]Node{}
	n[0] = Node{
		Source: 0,
		Dest:   1,
		Weight: 5,
	}
	n[1] = Node{
		Source: 1,
		Dest:   2,
		Weight: 4,
	}
	n[2] = Node{
		Source: 2,
		Dest:   3,
		Weight: 8,
	}
	n[3] = Node{
		Source: 3,
		Dest:   2,
		Weight: 8,
	}
	n[4] = Node{
		Source: 3,
		Dest:   4,
		Weight: 6,
	}
	n[5] = Node{
		Source: 0,
		Dest:   3,
		Weight: 5,
	}
	n[6] = Node{
		Source: 2,
		Dest:   4,
		Weight: 2,
	}
	n[7] = Node{
		Source: 4,
		Dest:   1,
		Weight: 3,
	}
	n[8] = Node{
		Source: 0,
		Dest:   4,
		Weight: 7,
	}

	return n
}
