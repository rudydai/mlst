package main

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
    if _, err := file.WriteString(string(len(e))+"\n"); err != nil {
        file.Close()
        return err
    }
    for num, edges := range(e) {
        _, err = file.WriteString(string(num)+"\n")
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
    var ret EdgeSet
    g := e.Graph()
    disjoint := make([]*Element, g.NumNodes)
    degree := make([]int, g.NumNodes)
    for i := 0; i < g.NumNodes; i++ {
        disjoint[i] = Makeset(i)
        degree[i] = 0
    }
    for node, adj := range(g.Neighbors) {
        connected := make(map[int]bool)
        newedges := 0
        for _, neighbor := range(adj) {
            neighborSet := Find(disjoint[neighbor])
            if neighborSet != Find(disjoint[node]) && !connected[neighborSet.value] {
                newedges += 1
                connected[neighborSet.value] = true
            }
        }
        if degree[node] + newedges >= 3 {
            for otherset := range(connected) {
                newEdge := Edge { [2]int{node, otherset} }
                newEdge.Normalize()
                ret[newEdge] = true
                Union(disjoint[node], disjoint[otherset])
                degree[node] += 1
                degree[otherset] += 1
            }
        }
    }
    return ret
}

func (e Edge) PrintForm() (s string) {
    return fmt.Sprintf("%d %d\n", e.Ends[0], e.Ends[1])
}

func start() {
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



