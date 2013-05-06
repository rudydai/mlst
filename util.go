package mlst

import (
    "os"
    "fmt"
)

//FILE INPUT AND OUTPUT
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

func (e Edge) PrintForm() (s string) {
    return fmt.Sprintf("%d %d\n", e.Ends[0], e.Ends[1])
}

//DISJOINT SETS
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
