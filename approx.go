package mlst

import (
    "os"
    "fmt"
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

func PrintSets(e []EdgeSet) (err error){
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
                Union(disjoint[node], disjoint[otherset])
                degree[node] += 1
                degree[neighbor] += 1
            }
        }
    }
    gout := ret.Graph()
    gout.Search()
    if gout.NumOfComponents == 1 {
        return ret
    }
    for node, origneighbors := range(g.Neighbors) {
        if degree[node] > 1 {
            for _, adj := range(origneighbors) {
                neighborset := Find(disjoint[adj])
                thisset := Find(disjoint[node])
                if thisset != neighborset && degree[adj] > 1 {
                    newedge := Edge { [2]int{node, adj} }
                    newedge.Normalize()
                    ret[newedge] = true
                    Union(thisset, neighborset)
                    degree[node] += 1
                    degree[adj] += 1
                }
            }
        }
    }
    gout.Search()
    if gout.NumOfComponents == 1 {
        return ret
    }
    for node, origneighbors := range(g.Neighbors) {
        if degree[node] > 0 {
            for _, adj := range(origneighbors) {
                neighborset := Find(disjoint[adj])
                thisset := Find(disjoint[node])
                if thisset != neighborset && degree[adj] > 0 {
                    newedge := Edge { [2]int{node, adj} }
                    newedge.Normalize()
                    ret[newedge] = true
                    Union(thisset, neighborset)
                    degree[node] += 1
                    degree[adj] += 1
                }
            }
        }
    }
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
            outsets[i] = ApproxSoln(edgeset)
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



