package main

import (
	"fmt"
	"log"
)

func printStruct(level Arguments, l int) {
    for n, lev := range level {
        s := *lev

        for x := 0; x < l; x++ {
            fmt.Print("  ")
        }
        fmt.Printf("%d: %s", n, s.Name)
        if len(s.Aliases) > 0 {
            fmt.Print(" [")
            for i, a := range s.Aliases {
                if i > 0 {
                    fmt.Print(", ")
                }
                fmt.Printf("%s", a)
            }
            fmt.Print("]")
        }
        fmt.Print("\n")
        if len(s.Children) > 0 {
            printStruct(s.Children, l+1)
        }
    }
}



func main() {

    structure, err := loadStructure("structure.json")
    if err != nil {
        log.Fatal(err)
    }
    printStruct(structure, 0)
}
