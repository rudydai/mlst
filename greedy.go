package mlst

import (
    "fmt"
)

func GreedySoln(e EdgeSet) (to_ret EdgeSet) {
	var answer EdgeSet = make(map[Edge]bool)
    g := e.Graph()
    //disjoint := make([]*Element, MaxNumNodes)
    //degree := make([]int, MaxNumNodes)

    var visitedNodes []int
    /**for edge := range e {
        fmt.Println("edge is %s", edge)
    }*/

    for len(visitedNodes) < g.NumNodes {
        max := 0
        var temp int

        if len(visitedNodes) == 0 {
            for node, attached := range g.Neighbors {
                if len(attached) > max {
                    max = len(attached)
                    temp = node
                }
            }
        } else {
            for node := range visitedNodes {
                if len(g.Neighbors[node]) > max {
                    max = len(g.Neighbors[node])
                    temp = node
                }
            }
        }

        //fmt.Println("visitedNodes is %q", visitedNodes)

        /*visited := false
        for node := range visitedNodes {
            if node == temp {
                visited = true
            }
        }
        if visited == false {
            visitedNodes = append(visitedNodes, temp)
        }*/

        if ContainsInt(visitedNodes, temp) == false {
            visitedNodes = append(visitedNodes, temp)
        }

        for neighbor := range g.Neighbors[temp] {
            if ContainsInt(visitedNodes, neighbor) == false {
                visitedNodes = append(visitedNodes, neighbor)
            }
            newEdge := Edge{[2]int{temp, neighbor}}
            newEdge.Normalize()
            answer[newEdge] = true
        }

        //fmt.Println("visitedNodes is %s", visitedNodes)
    }
    return answer
}

func GreedyStart() {
	edgesets := GetEdgeSets()
    if edgesets != nil {
        outsets := make([]EdgeSet, len(edgesets))
        for i, edgeset := range(edgesets) {
            outsets[i] = GreedySoln(edgeset)
        }
        err := PrintSets(outsets)
        if err != nil {
            fmt.Print("Problem writing output to file")
        }
    }
}
