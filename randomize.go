package mlst

import (
	"math/rand"
    //"fmt"
)

func ContainsInt(array []int, toFind int) (bool) {
    for i := 0; i < len(array); i++ {
        if array[i] == toFind {
            return true
        }
    }
    return false
}

func ShuffleAdjList(adjList *AdjList) () {
	for i := range *adjList {
	    j := rand.Intn(i + 1)
	    (*adjList)[i], (*adjList)[j] = (*adjList)[j], (*adjList)[i]
	}
}

func dfs(node int, mlst *EdgeSet, visited map[int]bool, adjList AdjList, g Graph) () {
	if visited[node] == true {
		return
	}

	visited[node] = true

	ShuffleAdjList(&g.Neighbors[node])
	
}

func RandomSoln(e EdgeSet) (solution EdgeSet) {
	//g := e.Graph()
	return nil
}

func RandomizeStart() {
	edgesets := GetEdgeSets()
    if edgesets != nil {
        outsets := make([]EdgeSet, len(edgesets))
        for i, edgeset := range(edgesets) {
            outsets[i] = ApproxSoln(edgeset)
        	//fmt.Println(outsets)
        }
        /*err := PrintSets(outsets)
        if err != nil {
            fmt.Print("Problem writing output to file")
        }*/
    }
}