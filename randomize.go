package mlst

import (
	"math/rand"
    "fmt"
)

func DFS(mlst EdgeSet, node int, visited map[int]bool, g *Graph) () {
	if _, ok := visited[node]; ok {
		return
	}

	visited[node] = true

	ShuffleAdjList(&(g.Neighbors[node]))

	connecting_nodes := (*g).Neighbors[node]

	for neighbor := range connecting_nodes {
		if _, ok := visited[neighbor]; ok == false {
			newEdge := Edge { [2]int{node, neighbor} }
            newEdge.Normalize()
            mlst[newEdge] = true
            DFS(mlst, neighbor, visited, g)
		}
	}
}

func NumLeaves(mlst EdgeSet) (count int) {
	numLeaves := make(map[int]int)
	for edge := range mlst {
		if _, ok := numLeaves[edge.Ends[0]]; ok {
			numLeaves[edge.Ends[0]] += 1
		} else {
			numLeaves[edge.Ends[0]] = 1
		}

		if _, ok := numLeaves[edge.Ends[1]]; ok {
			numLeaves[edge.Ends[1]] += 1
		} else {
			numLeaves[edge.Ends[1]] = 1
		}
	}

	count = 0
	for _, val := range numLeaves {
		if val == 1 {
			count++
		}
	}
	return count
}

func FindPath(mlst EdgeSet, node int, endNode int, visited map[int]bool) (path []Edge) {
	if _, ok := visited[node]; ok {
		return nil
	}

	visited[node] = true

	for edge := range mlst {
		if edge.Ends[0] == node && edge.Ends[1] == endNode {
			path = []Edge{edge}
			return path
		}

		recurseResult := FindPath(mlst, edge.Ends[1], endNode, visited)
		if edge.Ends[0] == node && recurseResult != nil {
			path = []Edge{edge}
			path = append(path, recurseResult...)
			return path
		}
	}
	return nil
}

func RandomSoln(e EdgeSet) (solution EdgeSet) {
	g := e.Graph()
	var mlst EdgeSet = make(map[Edge]bool)

	for true {
		//method 1
		for node, neighbors := range g.Neighbors {
			newEdge := Edge { [2]int{node, neighbors[rand.Intn(len(neighbors))]} }
            newEdge.Normalize()
            mlst[newEdge] = true
		}
		//checkIfMlst(mlst)

		//method 2
		mlst = make(map[Edge]bool)
		for i := 0; i < g.NumNodes - 1; {
			var randomEdge Edge
			for edge := range e {
				randomEdge = edge
				break
			}
			if _, ok := mlst[randomEdge]; !ok {
                mlst[randomEdge] = true
                i++
		    }
		}
		//checkIfMlst(mlst)

		visited := make(map[int]bool)
		mlst = make(map[Edge]bool)
		var startNode, endNode int
		for node, _ := range g.Neighbors {
			startNode = node
			DFS(mlst, node, visited, g)
			for tempNode, _ := range g.Neighbors {
				endNode = tempNode
				break
			}
			break
		}
		fmt.Println(mlst)

		fmt.Println(NumLeaves(mlst))
		fmt.Println(FindPath(mlst, startNode, endNode, visited))
		fmt.Println(FindPath(mlst, endNode, startNode, visited))
	}
	solution = mlst

	return solution
}

func RandomizeStart() {
	edgesets := GetEdgeSets()
    if edgesets != nil {
        outsets := make([]EdgeSet, len(edgesets))
        for i, edgeset := range(edgesets) {
            outsets[i] = ApproxSoln(edgeset)
            //fmt.Println(outsets)
        }
        err := PrintSets(outsets)
        if err != nil {
            fmt.Print("Problem writing output to file")
        }
    }
}
