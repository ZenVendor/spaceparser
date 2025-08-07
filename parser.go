package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Argument struct {
    Name string `json:",omitempty"`
    RequiresValue bool `json:",omitempty"`
    HasChildren bool `json:",omitempty"`
    RequiresChild bool `json:",omitempty"`
    MaxArgs int `json:",omitempty"`
    Aliases []string `json:",omitempty"`
    Children Arguments `json:",omitempty"`
}
type Arguments []*Argument

type Parser struct {
    args []string
    kwargs map[string]string
}


func loadStructure(path string) (Arguments, error) {
    var structure Arguments
    jsonData, err := os.ReadFile(path)
    if err != nil {
        return structure, fmt.Errorf("%w: %v", ErrReadingFile, path)
    }
    if err = json.Unmarshal(jsonData, &structure); err != nil {
        return structure, fmt.Errorf("%w: argument structure", ErrLoadingJson)
    }
    return structure, nil
}

