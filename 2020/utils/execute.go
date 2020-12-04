package utils

import (
	"time"
)

func Run(run func(string, bool) (int, int), input string, verbose bool) {
	start := time.Now()
	part1, part2 := run(input, verbose)
	elapsed := time.Since(start)

	Log(true, "info", "Run in", elapsed)
	Log(true, "warn", "Part 1:", part1, "\tPart 2:", part2)
}
