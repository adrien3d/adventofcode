package main

import (
	"github.com/adrien3d/adventofcode/2020/utils"
	"sort"
	"strconv"
	"strings"
)

func solve(input string, v bool) (part1TotalScore, part2TotalScore int) {
	lines := strings.Split(input, "\n")

	adapters := []int(nil)
	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		adapters = append(adapters, val)
	}
	sort.Slice(adapters, func(i, j int) bool { return adapters[i] < adapters[j] })
	utils.Log(v, "warn", adapters)

	var currentPos, oneDiff, threeDiff int
	for i := 0; i < len(adapters); i++ {
		if i == 0 {
			if adapters[0] == 1 {
				oneDiff++
				utils.Log(v, "info", "First: 1 diff")
			} else {
				threeDiff++
				utils.Log(v, "info", "First: 3 diff")
			}
		} else {
			if i == len(adapters)-1 { //End of array, no need to check next
				currentPos = currentPos - 1
			}
			if adapters[currentPos+1]-adapters[currentPos] == 1 {
				utils.Log(v, "info", "Regular: 1 diff from", adapters[currentPos], "to", adapters[currentPos+1])
				oneDiff++
			} else if adapters[currentPos+1]-adapters[currentPos] == 2 {
				utils.Log(v, "err", "2 diff")
				oneDiff += 2
			} else if adapters[currentPos+1]-adapters[currentPos] == 3 {
				utils.Log(v, "info", "Regular: 3 diff from", adapters[currentPos], "to", adapters[currentPos+1])
				threeDiff++
			} else {
				utils.Log(v, "err", "More than 3 diff")
			}
			currentPos++
		}
	}
	threeDiff++ // Connect to device
	utils.Log(v, "warn", oneDiff, threeDiff)

	arrangements := make([]int, len(adapters))
	arrangements[0]++
	for i := range adapters {
		for j := i + 1; j < len(adapters); j++ {
			if adapters[j]-adapters[i] > 3 {
				break
			}
			utils.Log(v, "warn", i, j, "step_size:", adapters[j]-adapters[i], "adding", arrangements[i], "to", arrangements[j], "arrangements:", arrangements)
			arrangements[j] += arrangements[i]
		}
	}
	part2 := arrangements[len(arrangements)-1]
	utils.Log(v, "err", arrangements, part2)

	return oneDiff * threeDiff, part2TotalScore
}

func main() {
	testInput := "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"
	testInput2 := "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"
	//realInput := "76\n51\n117\n97\n7\n77\n63\n18\n137\n10\n23\n14\n130\n131\n8\n91\n17\n29\n2\n36\n110\n35\n113\n30\n112\n61\n83\n122\n28\n75\n124\n82\n101\n135\n42\n44\n128\n32\n55\n85\n119\n114\n72\n111\n107\n123\n54\n3\n98\n96\n11\n62\n22\n49\n37\n1\n104\n43\n24\n31\n129\n69\n4\n21\n48\n39\n9\n38\n58\n125\n81\n89\n65\n90\n118\n64\n25\n138\n16\n78\n92\n102\n88\n95\n132\n47\n50\n15\n68\n84\n136\n103"
	utils.Run(solve, testInput, true)
	utils.Run(solve, testInput2, true)
	//utils.Run(solve, realInput, true)
}
