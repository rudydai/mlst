package mlst

import (
    "fmt"
)


func ApproxSoln(e EdgeSet) (to_ret EdgeSet) {
    var ret EdgeSet = make(map[Edge]bool)
    g := e.Graph()
    disjoint := make([]*Element, MaxNumNodes)
    degree := make([]int, MaxNumNodes)
    for i := 0; i < MaxNumNodes; i++ {
        disjoint[i] = Makeset(i)
        degree[i] = 0
    }
    for node, adj := range(g.Neighbors) {
        ownSet := Find(disjoint[node]).value
        connected := make(map[int]int)
        newedges := 0
        for _, neighbor := range(adj) {
            neighborSet := Find(disjoint[neighbor]).value
            if _, ok := connected[neighborSet]; !ok && neighborSet != ownSet {
                newedges += 1
                connected[neighborSet] = neighbor 
            }
        }
        if degree[node] + newedges >= 3 {
            for otherset, neighbor := range(connected) {
                newEdge := Edge { [2]int{node, neighbor} }
                newEdge.Normalize()
                ret[newEdge] = true
                //fmt.Printf("joining %d and %d\n", node, neighbor)
                Union(disjoint[node], disjoint[otherset])
                degree[node] += 1
                degree[neighbor] += 1
            }
        }
    }
    gout := ret.Graph()
    gout.Search()
    if gout.NumOfComponents == 1 && gout.NumNodes == g.NumNodes {
        return ret
    }
    //fmt.Print("now beginning connecting of leafy forest\n")
    for node, origneighbors := range(g.Neighbors) {
        if degree[node] > 1 {
            for _, adj := range(origneighbors) {
                thisset := Find(disjoint[node])
                neighborset := Find(disjoint[adj])
                if thisset != neighborset && degree[adj] > 1 {
                    newedge := Edge { [2]int{node, adj} }
                    newedge.Normalize()
                    ret[newedge] = true
                    //fmt.Printf("joining %d and %d\n", node, adj)
                    gout.AddEdge(newedge)
                    Union(thisset, neighborset)
                    degree[node] += 1
                    degree[adj] += 1
                }
            }
        }
    }
    gout.Search()
    if gout.NumOfComponents == 1 && gout.NumNodes == g.NumNodes {
        return ret
    }
    //fmt.Print("joining leaves as necessary\n")
    for node, origneighbors := range(g.Neighbors) {
        for _, adj := range(origneighbors) {
            neighborset := Find(disjoint[adj])
            thisset := Find(disjoint[node])
            if thisset != neighborset {
                newedge := Edge { [2]int{node, adj} }
                newedge.Normalize()
                ret[newedge] = true
                //fmt.Printf("joining %d and %d\n", node, adj)
                gout.AddEdge(newedge)
                Union(thisset, neighborset)
                degree[node] += 1
                degree[adj] += 1
            }
        }
    }
    /*gout.Search()
    if gout.NumOfComponents != 1 {
        fmt.Print("could not make a spanning tree\n")
    }*/
    return ret
}

func Start() {
    edgesets := GetEdgeSets()
    if edgesets != nil {
        outsets := make([]EdgeSet, len(edgesets))
        for i, edgeset := range(edgesets) {
            outsets[i] = ApproxSoln(edgeset)
        }
        err := PrintSets(outsets)
        if err != nil {
            fmt.Print("Problem writing output to file")
        }
    }
}

