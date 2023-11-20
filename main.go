package main

import (
	"fmt"
	"go-learning/common"
)

func removeDuplicate[T comparable](slice []T) []T {
	return common.ToSet[T](slice).ToSlice()
}

type Row struct {
	col1 string
	col2 string
}

func main() {
	rows := []Row{{col1: "k1", col2: "v1"}, {col1: "k2", col2: "v2"}, {col1: "k1", col2: "v1"}}
	rows = removeDuplicate[Row](rows)
	fmt.Println(rows)
}
