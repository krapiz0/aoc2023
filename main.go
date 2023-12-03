package main

import (
	"fmt"
	"time"

	// j"github.com/jqvk/aoc2023/day1"
	// "github.com/jqvk/aoc2023/day2"
	"github.com/jqvk/aoc2023/day3"
)

func measureAndPrint[T any](label string, fn func() T) {
	startTime := time.Now()
	res := fn()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("%s: %v\t\t(%v)\n", label, res, elapsed)
}

func main() {
	// measureAndPrint("D1P1", day1.Part1)
	// measureAndPrint("D1P2", day1.Part2)
	// measureAndPrint("D2P1", day2.Part1)
	// measureAndPrint("D2P2", day2.Part2)
	measureAndPrint("D3P1", day3.Part1)
}
