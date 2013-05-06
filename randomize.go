package mlst

import (
    //"fmt"
)

func dfs(node int, mlst EdgeSet, visited []bool)

func RandomSoln(e EdgeSet) (solution Edgeset) {
	g := e.Graph()
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