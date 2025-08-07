package main

import "errors"

var (
    ErrReadingFile = errors.New("Error reading file")
    ErrLoadingJson = errors.New("Error loading JSON data")

)
