package main

import (
	"github.com/adrien3d/adventofcode/2020/utils"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	argument int
}

func runInstructions(ref map[int]*instruction, verbose bool) int {
	var valuesAccessed []int
	var currentPos, accumulatorVal int
	for i := 0; i <= len(ref); i++ {
		if utils.IntArrayContains(valuesAccessed, currentPos) {
			utils.Log(verbose, "info", "Entering an infinite loop, reexecuting", currentPos, ref[currentPos], accumulatorVal)
			return accumulatorVal
		}
		valuesAccessed = append(valuesAccessed, currentPos)
		/*for j, val := range ref {
			fmt.Print(j, val)
		}
		fmt.Println()
		utils.Log(verbose, "warn", "run:", ref)*/
		//utils.Log(verbose, "warn", "Trying to access value:", currentPos, ref[currentPos], "of:", len(ref))
		if currentPos >= len(ref) {
			utils.Log(verbose, "err", "Trying to access value outside of ref, at:", currentPos, "len(ref)=", len(ref))
			break
		}
		switch ref[currentPos].operation {
		case "acc":
			accumulatorVal += ref[currentPos].argument
			currentPos++
		case "jmp":
			currentPos += ref[currentPos].argument
		case "nop":
			currentPos++
		}
	}
	return -1
}

func solve(input string, verbose bool) (part1TotalScore, part2TotalScore int) {
	lines := strings.Split(input, "\n")

	ref := make(map[int]*instruction)
	for i, line := range lines {
		inst := strings.Split(line, " ")
		arg, _ := strconv.Atoi(inst[1])
		ref[i] = &instruction{inst[0],arg}
	}
	utils.Log(verbose, "err", len(ref))
	//replace one instruction from jmp to nop, or nop to jmp
	for i := 0; i < len(ref); i++ {
		refModified := make(map[int]*instruction)
		for j, _ := range ref {
			refModified[j] = ref[j]
		}
		switch refModified[i].operation {
		case "acc":
			continue
		case "jmp":
			utils.Log(verbose, "warn", "Changing instruction", ref[i], "from jmp to nop")
			refModified[i].operation = "nop"
			//refModified[i].argument = ref[i].argument
			runInstructions(refModified, verbose)
		case "nop":
			if refModified[i].argument != 0 {
				utils.Log(verbose, "warn", "Changing instruction", ref[i], "from nop to jmp")
				refModified[i].operation = "jmp"
				//refModified[i].argument = ref[i].argument
				runInstructions(refModified, verbose)
			}
		}
	}

	return runInstructions(ref, true), part2TotalScore
}

func main() {
	testInput := "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"
	realInput := "acc +6\nacc +21\nnop +297\njmp +2\njmp +156\nacc -7\nacc +6\nacc -16\nacc +14\njmp +127\nacc +8\nacc -16\nacc +48\nnop +189\njmp +214\nacc +20\nacc +25\nacc +3\nacc +10\njmp +571\nacc -7\nacc -16\nacc +29\nacc +24\njmp +346\nacc +1\nacc -11\nacc -14\nacc +14\njmp +18\nacc -5\nacc +22\njmp +351\nacc +13\nacc +6\nnop -2\nacc +20\njmp +405\nacc +13\nacc +40\njmp +265\nacc +32\nacc -3\nacc +13\nacc +9\njmp +175\njmp +237\nnop +113\njmp +127\nacc -10\nacc +49\nnop -36\nacc +17\njmp +156\nnop +181\njmp +419\nacc +11\njmp +1\nacc +49\nnop +187\njmp +427\nnop +568\nacc -14\nnop +481\nacc +47\njmp +163\njmp +309\njmp +410\nacc +23\nacc -17\nnop +229\nacc -4\njmp +538\nacc -12\njmp +200\nacc +35\nacc +22\nacc -14\nacc +46\njmp +388\nacc +20\nacc +15\nacc +0\nacc +26\njmp +197\njmp +267\nacc +33\njmp -82\nacc +44\nacc -19\njmp +42\nacc +14\njmp +1\njmp +263\nacc -18\njmp +80\nacc -12\njmp +248\nacc +0\nacc +44\nacc +6\njmp +134\njmp +1\nacc +2\nacc +38\njmp +40\nacc +30\nacc -1\njmp +141\njmp +1\njmp +460\nnop +241\nacc +34\nacc +35\njmp +1\njmp -6\nnop +172\nacc +6\nnop +299\njmp +298\nacc +20\nnop +52\nacc +47\nacc +32\njmp +293\nacc +22\nacc -4\njmp +417\nacc -8\njmp +276\njmp +452\njmp +381\nacc +0\nnop -79\njmp +101\njmp +253\nacc +30\nacc +42\nacc +21\njmp +231\nacc +35\nacc +20\nacc -10\nacc -19\njmp +173\njmp +2\njmp -16\nacc -8\nacc +47\nacc -9\nacc +11\njmp -35\nacc -2\njmp +121\nacc -16\nacc -19\nacc +47\nacc +34\njmp -11\njmp +34\nacc +40\nacc -3\nacc -18\nnop +359\njmp +9\nacc +26\njmp +117\nacc -8\nacc -14\njmp +1\njmp +123\nnop +117\njmp -63\njmp +13\njmp +227\nacc +41\njmp +207\nnop -86\nacc +37\nacc -15\njmp +55\nacc +24\nacc +8\njmp +1\njmp +79\njmp +270\njmp -125\njmp +214\njmp +100\nacc -8\nacc -19\nacc +23\njmp -167\nnop -18\nacc -14\nacc +19\nnop +291\njmp +361\nacc +25\nacc +21\nnop +99\njmp +229\nnop +228\nacc +4\nacc +24\njmp -12\njmp +1\nacc +26\njmp +75\nacc +22\nnop +105\nacc +46\nacc -8\njmp -81\nacc +46\njmp -168\nacc +14\njmp -57\nacc -13\njmp +137\nnop +362\nacc +28\njmp +352\nacc +8\nacc +21\nacc +30\nacc +13\njmp -91\nacc +39\njmp +296\nacc +27\nacc +16\nacc +5\nacc +48\njmp -6\nnop -210\nacc +29\nacc +47\njmp -78\njmp +228\nacc +13\nacc -11\nacc +3\njmp +96\nacc +0\njmp +313\nacc +30\njmp +251\njmp +203\nnop -202\nnop -177\nacc -17\nacc +30\njmp -128\njmp +227\nnop +84\nacc +21\nacc +3\njmp -18\nacc +33\nnop -128\nnop +368\njmp -87\nacc +30\njmp +88\nacc -3\nacc +17\njmp +63\nacc +37\nacc -13\njmp +340\njmp +1\nacc +11\njmp +325\njmp -58\nacc +43\njmp +23\njmp +157\nacc +35\nacc +10\njmp +25\njmp +124\njmp -109\nnop +40\njmp +183\nacc +46\nacc +37\njmp +88\nacc -8\njmp +162\nacc +4\nacc +22\njmp +220\nacc +32\njmp -214\nacc +3\nnop -56\nacc +30\njmp -138\nacc +22\njmp +329\nacc +12\nnop +115\nacc +38\njmp -231\nacc +9\njmp +1\nacc +25\nacc +47\njmp +329\nacc +14\nnop +220\nacc -19\nacc -19\njmp +225\njmp -248\njmp +249\njmp -231\nnop -30\nacc +26\nacc +32\nacc +12\njmp +190\nacc +4\njmp -251\nacc +20\nnop -27\nacc +20\nacc +16\njmp -41\nacc +21\nacc +45\nacc +24\njmp -65\nacc +39\nacc -9\nacc -12\nacc +23\njmp -50\nacc +49\nnop -105\nacc +17\njmp +180\nacc +14\njmp +238\nacc +1\nnop -285\nacc +26\nacc -15\njmp +216\nnop -95\njmp +60\njmp -261\nacc +49\nacc +31\njmp +210\njmp +1\nacc -5\njmp -304\nacc +48\nacc +0\nnop +2\njmp -347\nacc +50\njmp +148\nacc +5\nnop -193\nacc -5\nacc +0\njmp +221\nacc +39\nacc -14\njmp +39\nacc +24\njmp -100\nacc +0\nacc -16\nacc +4\njmp -190\nacc +21\nacc -2\nacc -16\njmp +162\njmp +28\nacc +26\nacc +19\nacc -5\njmp -362\nacc -16\nnop -6\njmp -223\njmp +101\nacc -7\nacc -10\nacc -16\njmp -146\nnop +126\nacc -18\nacc +25\nnop -232\njmp +61\njmp -86\nacc -9\nnop -20\njmp -318\njmp -90\nnop -315\nacc +33\nacc +0\nacc +18\njmp +41\njmp +1\njmp +8\nacc +4\nacc +13\nacc -19\njmp -128\nacc +46\njmp -322\njmp -404\nacc -3\nacc +24\nacc +5\nnop +159\njmp -322\nacc -10\nacc -18\nacc +42\nacc +44\njmp -130\nacc -4\nacc +4\njmp -302\njmp +18\nacc -3\nacc -18\nacc +29\nacc +36\njmp -405\nacc +15\nnop +91\njmp +98\nacc -14\nnop +88\nacc +16\njmp +86\nacc -14\nacc +20\nacc -10\nnop -7\njmp +81\njmp +74\nacc +3\nacc +18\njmp -294\nacc +16\njmp -173\nnop +65\nacc +31\nacc -14\njmp -77\nacc +16\nnop -415\njmp -33\nacc +41\nnop -398\njmp +46\nacc +50\nacc +7\nacc -19\nnop -15\njmp -317\nacc +38\njmp -298\nacc -3\njmp -273\njmp -120\nnop +10\nacc +26\njmp +39\njmp -348\nnop -204\nacc +0\nacc +0\nnop -179\njmp -208\njmp -52\nnop -299\nacc +36\nacc +14\njmp +108\njmp -476\njmp -283\nacc +31\nacc +33\nacc +32\nacc -13\njmp -295\nacc +42\nacc +27\njmp -112\nacc +37\nacc +22\nacc +34\njmp -236\nacc -2\nacc +33\nnop +75\nnop -503\njmp -146\nnop -324\njmp -148\nacc +39\njmp -112\nacc +32\nacc -15\nacc -5\njmp +8\nacc +30\njmp -196\nacc +5\njmp -361\nacc +29\njmp -153\njmp -26\njmp -344\nacc +39\nacc +25\nnop -153\nacc +34\njmp -79\nnop -97\nnop -90\nacc +21\nacc -16\njmp +4\nacc +41\njmp +1\njmp -233\nacc +37\njmp -235\nacc +41\nacc +44\njmp -12\nnop +1\nacc -18\njmp +74\nacc +49\nnop -30\nacc -17\njmp -364\nacc -13\nacc +50\njmp -223\nnop -507\nacc +17\nacc +2\njmp -302\nacc +29\njmp -8\nacc +33\nnop -514\nacc +15\nacc +31\njmp -189\nnop -98\nacc +47\nacc +21\njmp -322\njmp -77\njmp -555\nacc +44\nnop -142\njmp -96\nacc +45\nacc +44\nnop -121\njmp -322\nacc +35\nacc +48\nacc +0\njmp -393\njmp +1\nacc +0\nnop -356\njmp -512\nacc +31\nacc -10\nacc +37\njmp -71\nacc +25\nacc -8\njmp -200\nacc -7\nacc -12\nacc +0\nacc -16\njmp -369\nacc +14\nacc +11\nacc +35\njmp -510\njmp -472\nacc +39\nacc +27\nacc +50\nnop -360\njmp -498\nacc +13\njmp -500\nacc -17\nacc +11\nacc +10\nacc +7\njmp -11\nacc +6\nacc +30\nacc +25\nacc -19\njmp -373\nacc -4\nacc +19\nnop -329\njmp -582\nacc +11\nacc +26\nacc -1\nacc +35\njmp -548\nnop -59\nnop -576\nacc +27\njmp -145\nacc -7\nacc +41\nnop -272\nacc +35\njmp -170\nacc -6\nacc -4\nacc +33\nacc +33\njmp +1"
	utils.Run(solve, testInput, true)
	utils.Run(solve, realInput, true)
}
