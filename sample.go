package sample2

import (
	"strings"

	"github.com/keep94/gochart"
)

func Squares() string {
	xs := gochart.NewInts(1, 1, 10)
	ys := xs.Apply(func(x int64) int64 { return x * x })
	var builder strings.Builder
	gochart.NewChart(xs, ys).WriteTo(&builder)
	return builder.String()
}
