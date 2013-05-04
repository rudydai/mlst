package mlst

import (
    "os"
    "fmt"
    "math/rand"
)

func GetEdgeSets() (e []EdgeSet) {
    var infile string
    if len(args) < 1 {
        infile = DefaultInputFile
    } else {
        infile = args[0]
    }
    fmt.Print(infile)
    file, err := os.Open(infile)
    if err != nil {
        printError(fmt.Sprintf("Cannot open file"))
        file.Close()
        return nil
    }
    inReader := NewInFileReader(file)
    if edgeSets, err := inReader.ReadInputFile(); err != nil {
        fmt.Print("Invalid format of infile")
    }else {
        file.Close()
        return edgeSets
    }
    file.Close()
    return nil
}

func PrintSets(e []EdgeSet) (err error) {
    var outfile string
    if len(args) < 2 {
        outfile = DefaultOutputFile
    } else {
        outfile = args[1]
    }
    file, err := os.Create(outfile)
    if err != nil {
        fmt.Print("could not create file for writing")
        file.Close()
        return err
    }
    if _, err := file.WriteString(fmt.Sprintf("%d\n", len(e))); err != nil {
        file.Close()
        return err
    }
    for _, edges := range(e) {
        _, err = file.WriteString(fmt.Sprintf("%d\n", len(edges)))
        if err != nil {
            file.Close()
            return err
        }
        for edge := range(edges) {
            _, err = file.WriteString(edge.PrintForm())
            if err != nil {
                file.Close()
                return err
            }
        }
    }
    return nil
}

func ManyIters(e EdgeSet) (to_ret EdgeSet) {
    var BRUTE_ITERS int = 1000
    var best EdgeSet = nil
    var curr EdgeSet
    var numLeaves
    var mostLeaves int = 0
    for i := 0; i < BRUTE_ITERS; i++ {
        ordering := rand.Perm(MaxNumNodes)
        curr = ApproxSoln(e, ordering)
        currg = curr.Graph()
        currg.Search()
        numLeaves = currg.NumLeaves
        if numLeaves > mostLeaves {
            mostLeaves = numLeaves
            best = curr
        }
    }
    return best
}

func ApproxSoln(e EdgeSet, order []int ) (to_ret EdgeSet) {
    var ret EdgeSet = make(map[Edge]bool)
    g := e.Graph()
    disjoint := make([]*Element, MaxNumNodes)
    degree := make([]int, MaxNumNodes)
    for i := 0; i < MaxNumNodes; i++ {
        disjoint[i] = Makeset(i)
        degree[i] = 0
    }
    for _, node := range(order) {
        adj := g.Neighbors[node]
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


func (e Edge) PrintForm() (s string) {
    return fmt.Sprintf("%d %d\n", e.Ends[0], e.Ends[1])
}

func Start() {
    edgesets := GetEdgeSets()
    if edgesets != nil {
        outsets := make([]EdgeSet, len(edgesets))
        for i, edgeset := range(edgesets) {
            outsets[i] = ManyIters(edgeset)
        }
        err := PrintSets(outsets)
        if err != nil {
            fmt.Print("Problem writing output to file")
        }
    }
}

type Element struct {
    Parent *Element
    value int
}

func Makeset(val int) (*Element) {
    e := new(Element)
    e.Parent = e
    e.value = val
    return e
}

func Find(e *Element) (*Element) {
    if e.Parent == e {
        return e
    }
    e.Parent = Find(e.Parent);
    return e.Parent
}

func Union(e1,e2 *Element) () {
    root1 := Find(e1)
    root2 := Find(e2)
    root1.Parent = root2
}



