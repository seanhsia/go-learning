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
	col3 int
}

func main() {
	rows := []Row{{col1: "k1", col2: "v1", col3: 10}, {col1: "k2", col2: "v2", col3: 12}, {col1: "k1", col2: "v1", col3: 10}, {col1: "k3", col2: "v3", col3: 0}}
	rows = removeDuplicate(rows)
	vals := common.Map(rows, func(r Row) int { return r.col3 })
	vals = common.Filter(vals, func(i int) bool { return i > 0 })
	sum := common.Reduce(vals, 0, func(i1, i2 int) int { return i1 + i2 })
	fmt.Println(sum)
}
