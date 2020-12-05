package main

import (
	"fmt"
	"github.com/adrien3d/adventofcode/2020/utils"
	"strings"
)

type PasswordLine struct {
	min, max int
	char     byte
	password string
}

func ParseCustomList(input, separator string) []PasswordLine {
	lines := strings.Split(input, separator)
	list := make([]PasswordLine, len(lines))
	for i, line := range lines {
		var lineData PasswordLine
		_, err := fmt.Sscanf(line, "%d-%d %c: %s", &lineData.min, &lineData.max, &lineData.char, &lineData.password)
		utils.CheckErr(err)
		list[i] = lineData
	}
	return list
}

func day2part1(input string) {
	res := 0
	ret := ParseCustomList(input, "\n")
	for _, pwdLine := range ret {
		occurences := 0
		for _, character := range pwdLine.password {
			if string(character) == string(pwdLine.char) {
				occurences++
			}
		}
		if occurences >= pwdLine.min && occurences <= pwdLine.max {
			res++
		}
		//utils.Log("info", pwdLine, occurences)
	}
	utils.Log(true, "warn", res, "lines satisfies part 1 requirements")
}

func day2part2(input string) {
	res := 0
	ret := ParseCustomList(input, "\n")
	for _, pwdLine := range ret {
		occurences := 0
		for i, character := range pwdLine.password {
			if (string(character) == string(pwdLine.char)) && (pwdLine.min == i+1 || pwdLine.max == i+1) {
				occurences++
			}
		}
		if occurences == 1 {
			res++
		}
		//utils.Log("info", pwdLine, occurences)
	}
	utils.Log(true, "warn", res, "lines satisfies part 2 requirements")
}

func main() {
	testInput := "1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc"
	realInput := "2-5 z: zzztvz\n2-8 d: pddzddkdvqgxndd\n4-14 r: rrrjrrrrrrbrrccrr\n2-7 r: zrgsnrr\n9-10 z: zzzxwzznpd\n8-13 g: gggggggxgggghggg\n1-6 c: xcccxcccccz\n3-4 b: bxbt\n8-11 d: dddddddzddv\n4-14 m: kxdmmmdmfwmmmdfr\n16-17 g: ggggggggggggggggg\n2-3 f: ddvdlff\n9-11 g: ggggpfgggggg\n7-8 c: fdhctccc\n3-6 c: tmcdcncqcvccg\n2-3 l: clllll\n1-9 b: bbbbbbbbbb\n10-15 w: wglmwwwrnnzgwhhwvvd\n10-14 g: ggggggxsgggqggg\n9-19 q: fjlqbvtdngwvtbnsgfm\n1-2 k: kgbkkv\n4-10 m: mmmmtmmmmgm\n3-4 h: hplh\n5-7 q: qqqqrqk\n3-12 v: zbvlbpxcrnvvwjpwl\n1-4 r: crrr\n4-5 h: hhhrj\n2-4 r: mhrr\n9-10 c: pccccsccbk\n10-16 n: drnhqnnnnqnnpxmh\n5-6 v: vvvnvvg\n11-12 c: ccccgcccccpc\n5-10 b: bdbjbbfmfbb\n9-13 w: zwbwwzhwwkwsxdfwglx\n5-19 v: vvvvvqrvvfvgvvvvmvv\n3-9 d: dddgztdljb\n12-13 s: swssssrsssshx\n5-6 l: dllkljq\n4-5 q: qqqsdqq\n2-4 w: tmwg\n10-12 l: llllsldflllllljjlxl\n3-11 s: sjtsssssksbs\n1-5 v: hvvtmvgvv\n3-12 g: ggvgkgfgggghgg\n3-4 p: ccpp\n3-9 v: vfvvvvvvvvvvvr\n8-16 k: kkkkkkkkkkkkkwkkkk\n8-9 x: wxbrxfdrw\n7-9 n: nnnfnnnnn\n7-11 w: wwwwxwwwwwzxw\n6-9 n: nknnhnnnnlnnn\n10-12 l: qllllllllqbl\n17-19 p: pppppppppppppwppkpp\n4-5 q: qqqqqq\n7-15 v: vvvbmbjvjznzdhfvtn\n3-12 w: wwwwwdwwwwwwtwwwww\n7-10 x: mmgmknpvdbxdl\n9-10 v: jljttxpvvv\n11-15 d: dshddnddqdhddvm\n1-3 r: vrqrrr\n6-7 c: ltcbpccw\n5-13 q: qqqqcqqqxqrqg\n3-7 l: lhlmhzl\n2-3 l: lxntjnttkllll\n3-4 b: bbnsb\n5-6 j: rjppzvjdjj\n4-10 g: fgggjgvfggggggsggm\n11-12 z: ztzzzkzzzzsvzzzzt\n2-3 h: hcvbf\n7-13 s: sswkwhsdqsdcs\n3-10 j: mghjglnbdhjp\n6-8 x: dfxrxxsxpsxxxxmc\n11-12 f: fffmfcdpffmpflfffdw\n4-7 n: nnnknncnn\n2-6 k: tzqlzqhrncvktfwvswt\n2-4 h: pqplh\n8-13 h: zhlwqvmnkjhhtkf\n2-3 l: llll\n19-20 j: jjjjjjjjjjjjjjjjjrjj\n3-14 n: crnnrgpmxxlrfljk\n2-4 t: stpr\n11-13 c: cctcccccccccccc\n5-6 q: fqqqqqdwtwqjq\n6-7 z: zzzzzzg\n8-16 j: njjgsjjljktjcbjvjc\n3-6 c: swjccc\n5-7 h: ghvwjmzm\n9-12 s: sssssmvtdckmfs\n3-4 r: bsrm\n6-7 c: ccctcccc\n6-10 t: swndrtcvtl\n13-19 v: vjvxvvvdvvvvvvdvblv\n5-7 w: hnwwwtwp\n8-9 k: kkkkkktzxkkv\n4-6 g: xpsxgs\n2-3 r: xrrdxm\n2-9 f: ffvqxxvsf\n7-10 t: tttttttttt\n4-6 q: bbqqqqqqbq\n7-9 q: xjcwqhprgh\n3-6 j: jjjjjjjxjjj\n1-3 l: blll\n2-5 p: phbppp\n3-4 p: phpb\n2-15 r: mwgnrkxrrbmgpzrdrr\n8-13 x: jsrlvxmxwglkh\n4-11 b: ctcbsjgjplbthhrskfnb\n6-7 q: qqqqqfvcq\n10-11 d: ddddddjdbndddd\n8-9 p: pppppppkf\n10-14 q: qqqqqqnqzgqsqv\n12-13 b: bbbbbbbbbbbbb\n6-11 h: khnkthfhhdhrhh\n5-6 g: gzggwcgq\n2-4 f: fbfg\n6-9 f: ffvfsfvffxff\n6-7 t: tsfdltt\n14-17 t: tttttttttttxtcttp\n3-4 v: vwtv\n7-10 x: pxmldqqxmw\n1-2 s: pdssssssss\n5-6 n: snwnml\n3-8 t: tjnpjqzxbj\n5-7 n: mxkqfnjb\n4-12 q: qqqkqqqqqqqkqqq\n4-5 d: pjhnsfdgddjsddvpxxd\n14-15 t: trttdtxwtdmtttt\n3-4 j: jjvz\n1-2 p: xqpppp\n11-15 d: dddpdddddddddddq\n15-16 b: bnbbbbbbbbbtbbdbb\n9-11 w: wwwnwwmwbghw\n15-18 g: tjwqbgwmvpvghxgzkgj\n3-4 j: jjps\n9-12 f: ffflfffmfnffvff\n1-3 g: gbgstndgggqrdbtpdqdr\n4-6 f: fflxfqffff\n10-14 t: ttttttttttttttttt\n1-7 v: vvsqlvvvvkj\n3-7 b: wvbxlwbjtbbbvh\n3-7 p: hqvsvpn\n2-6 f: fkwfhf\n3-11 d: dzrwnffrdvs\n8-9 s: sssdssztlssqksx\n12-15 t: tttttttttthtttt\n5-6 m: mmmmmmmm\n2-3 n: dnnl\n12-16 w: wrwwtvwwwwwwwwww\n2-4 j: jjjjsj\n7-11 x: xxxxxtcxxxlxxx\n8-17 n: nbngnnnsnjnbhnntzwxt\n5-6 b: tmbdgn\n15-16 x: xqlpjxbbxnrbxxxxqmkf\n1-3 c: ccclf\n12-14 m: zhmmmzltwhrmfmx\n15-17 c: ccvcxcmccgccccbcccc\n10-13 k: mkgkljkkkjkkg\n1-4 b: ckdr\n10-12 b: bvfbgbbbbbbb\n2-4 v: jfvcv\n2-7 p: jswmqpmc\n11-16 b: bbbncbpmbjbbvgrb\n1-6 n: ktjcwr\n2-9 l: lllrwljlnlwllln\n17-18 h: hhhhhhhhhhhhhhhhzfh\n9-14 b: bbbsbqbbjjbbkb\n7-12 j: jjjjjwwjjjjcj\n10-13 w: wfwwwwwwwjwww\n2-6 j: bjghpj\n13-16 z: zzzzzzzzbzzzzzvg\n10-11 f: ffffkfffcfxff\n3-4 m: mmmmkjvzmkp\n2-5 k: kkkkk\n5-6 r: rrbrrs\n1-2 c: ccqjmc\n13-14 r: rrrrrrrrrrrrztr\n8-12 z: zrzzzzzqbzvdz\n4-9 m: hlgmtvvqs\n9-10 t: tthttrtbrnt\n2-4 k: snpkzhck\n12-13 x: wlxwwxxxhgxxx\n11-16 h: ncmfhhvgghhhhhjhd\n13-15 g: sgxgfntgjnxkgbgm\n4-13 s: vqkslldmgmfrd\n3-14 z: zzzzzzzzzzzzzzzz\n14-15 c: ccccvccccccsmlpcfccx\n4-5 p: pppppp\n9-10 z: kgblzrqbzz\n13-15 p: pppppptpmppbpppqpp\n7-8 b: dntbbcmswb\n16-17 d: dddddddddddddddgd\n10-11 r: rrrsrmrrrqtr\n2-3 q: qtqgz\n6-8 c: xztzfkcch\n14-15 l: lllflllllllllzsl\n1-6 s: zssssss\n5-8 d: hdddwddvdpd\n10-12 j: jjjjjjjjjpjbj\n1-9 q: tqqqqqqqqnqqwqk\n2-4 t: shpwwttwqb\n8-9 s: xrssssssz\n3-4 c: cjcfc\n3-10 v: rfvzsncmfv\n4-7 t: lshczltjfqkmkt\n13-16 x: xxxxxxxfxxxxzxxrxx\n1-3 z: pzkzlzmzz\n4-5 d: ddkbfdtdd\n9-10 r: rrrrrrrrlfkj\n3-7 s: svtbsssjsldts\n10-13 c: ccccgxzqtctrc\n1-7 q: zgwqgrx\n6-8 f: ffnffrfg\n4-5 f: kffcffp\n1-2 m: ggcm\n5-6 z: zzzdgz\n9-14 j: jjjnjjjjjjjjjjjj\n2-17 d: xttgzggnfhsrjqdbppps\n5-6 f: xrcfff\n1-3 z: wzzg\n1-2 h: hhhvl\n13-17 q: qqqqqqqqqqmqqqqqqq\n2-3 q: bjqq\n1-5 w: zwwwgww\n5-7 s: jnvpmgj\n9-10 q: zctgztwpqq\n16-20 j: jkjjkjnjwqqjjjzgjjpk\n16-20 k: kkkkkkkkkkkkkkkkkkkk\n17-18 d: dddddddddddddddddbd\n3-5 d: dngrhxb\n10-12 t: tmttttttttrttwh\n7-9 q: nqqqkqcqqqqq\n7-9 z: zqzzzzwctng\n10-13 d: dddjldsddkddjpddd\n3-7 g: lgglgggggphx\n6-7 b: llkbcdxwtzrblp\n1-5 s: dssdscpzwcn\n2-6 h: hhphhhhhh\n4-13 z: smkzzpzjfzzbzqzxzjz\n2-5 j: jjjjjj\n8-11 p: lgmppspppppk\n9-10 c: ccccccccdnkc\n3-9 m: mmfmmmmpkmsmmm\n13-15 m: mmmmkmmmmmmmgmhmmm\n5-6 x: xxxxxxx\n10-14 b: bbbjbqbbwcbbbb\n3-12 f: ffmfffffffnzfffff\n11-12 m: mmmnmmmmmmmmm\n8-12 v: jptmbccvmvsxchfvv\n4-5 s: ssssns\n12-14 j: jjjjjjdjjjjvjhjjj\n11-13 j: wjjjljjjcjjjjqjjp\n1-7 g: ggggggggdg\n2-3 v: bkjqv\n4-13 q: qzhldwthbwgtr\n12-13 p: ppplppdgpppppp\n11-12 f: flffxfpffprgfff\n8-14 b: mbbbbbtbwbhstbbbbbbb\n6-7 m: bqmsmmm\n1-3 w: qxlggmxtwj\n3-17 c: dwbkjcwqcrccbcxcnvxc\n11-13 p: ppxppppdppxpzpppp\n4-10 v: vfvvfvzvqv\n1-7 m: mmmxmmmmmbmmrmmmmmm\n2-12 x: xrdxwxzwfxpxtqxxxgxp\n6-11 w: swzwwkddmsrxwzfg\n6-8 g: ggggggggggg\n3-4 h: tzshs\n7-8 t: zjtzpclm\n5-6 k: brzkzk\n14-15 c: cccccccccjccccc\n2-3 h: hhfh\n15-19 w: wwwwwwwkzxwwwwwwwwh\n8-10 f: mjmwqnlnhf\n3-8 l: ksllllplllc\n1-3 q: gjmhq\n2-9 l: lwklcxllf\n5-7 p: ppcclscppgdxmg\n3-6 z: fpxmpz\n8-11 r: vrrrrrrrrrrrrrs\n1-7 q: cqqqqqgq\n13-14 t: tttztttttttttt\n4-6 q: qzkqsqzmhhq\n8-14 l: llllfllplllgzll\n5-9 m: mmmtmmmmm\n9-14 l: llllllllrllllltl\n3-4 r: czrr\n2-5 s: lssslssnmfqss\n14-17 h: hhhhhhhhhhhhhhhhhh\n3-4 x: xxxx\n7-8 x: zdjbxxsl\n10-13 n: ffrnphpzhnfxnncptmn\n7-11 m: mmmmmmdmmmmmf\n1-2 j: jjjjjjjjjtj\n2-6 g: kglgsq\n3-8 x: csvxxxxxztxdxpw\n3-9 m: qbqwfmfpjmbnbcccgfgk\n2-14 v: rmqhqlrvcjdhpw\n8-9 p: pppppppvwpm\n8-12 f: zvfjxbdgzjbplm\n2-3 f: frff\n7-9 f: htkxffclfr\n4-7 v: cvqjctzvzvvv\n6-7 f: fjffflj\n1-4 b: tbnpb\n8-10 t: tjpbtfzhtfsdjtlttttt\n8-12 h: hhhhhhhhhhhrh\n3-5 n: hxbwgjchklnplnb\n1-2 z: jzvz\n4-7 z: vzzzvcz\n11-12 f: fffffvcfskcfcfftqf\n2-5 p: cxktpckkqkgjpptvgp\n9-16 b: bbbbbbbbkbbbbbbbbbbb\n10-11 p: jcpnpppxpppb\n7-9 n: sncnrzncnzdnngn\n10-11 n: ttqrrkbspnn\n11-19 c: ccczncccccsccpccccjc\n1-2 h: hhhbwhhknzf\n5-8 q: zxbtktkf\n2-6 x: trxhzl\n1-5 d: dskddkqdgqkhcz\n7-12 c: ccccccccccccc\n7-9 t: kskrrhtxtvsqzt\n5-12 v: vvvjvwrdfscv\n6-10 z: pzznznzczkwzzzmzz\n15-18 t: tttttntttttttttttttt\n4-6 m: bxmmmm\n2-8 k: tkkhqwkkkpk\n6-11 p: pppppgmpppppp\n7-10 b: bkbcbbrrbr\n6-11 c: cctccccbnbcvccc\n2-5 t: ccpxgwzqktxwjf\n4-11 s: sssssscsssssg\n3-4 d: ndpk\n3-4 z: gjzzzkbhzzs\n8-9 t: ttttttttt\n5-10 l: lllldmmlln\n9-15 j: njzjgjjjjjjjprjjj\n5-7 p: zppppppgkp\n2-6 r: rvdrrrgdnrrrlrr\n6-7 k: kkkkkkkk\n3-4 x: nxxxmtmljgldjx\n2-6 f: qfxxxftf\n1-9 z: zzdlqgzwzzjw\n6-8 n: rnjnstwm\n4-6 w: vhfhwckw\n3-7 k: kkpkkfpk\n5-7 k: kkkfzkhkk\n1-11 j: jjjtjjjjmjs\n2-17 w: wlwdkkwmkwskpstxvrwr\n5-11 g: kgpgjspqlwnx\n2-6 l: lpllld\n1-6 d: xwkvgvkdxlljpdqwmhs\n5-6 z: llrzzsz\n4-5 l: lllll\n18-19 g: gggggggngggggggggggg\n1-10 x: kkdxwkrtfh\n3-4 t: ttrqzttbsrt\n14-16 s: ssshsssssssssnsss\n2-6 k: kjzkxkkhkp\n2-10 b: bbbbbbbbbbbbbb\n1-9 v: vvvvvvvvvvv\n11-12 m: mqmmmmmmmmhsl\n5-6 s: sssmsss\n14-20 s: srssssssdsssskssksvs\n1-3 r: rrrr\n4-7 d: zddkdptdfdw\n6-7 g: rggglggtgkgg\n3-5 c: bcctcvpcprcccqccpcrc\n8-9 r: rrrxzrrlmrd\n12-17 c: hcctccccbqcctscsscv\n15-16 d: dkddddjdddddsdhwdddj\n13-14 b: blbbbbqbhbbbmz\n2-10 p: hvxqpjqstv\n1-4 p: plxp\n3-6 z: zzdlzqjzzz\n1-14 s: qsssssssslshswsjh\n3-5 r: grrrnrrrsvrrrdrrrjrr\n1-4 d: dfdjdxbpdddf\n2-3 p: pppppp\n16-18 c: cmcccjxfclcgmchswq\n9-11 x: xxrxxxxxvxrxx\n13-14 p: dpppppsppspftpppwp\n15-19 h: hhhhhhhhhhhhhhhhhhhh\n6-12 q: qqqxqqqqgqqpqqqqqqqv\n10-12 g: gqgdgzhgghjgdd\n2-6 p: pnppppp\n3-14 l: lllllllllllllllll\n12-19 h: hhdhhhhvxrfhcshhwxh\n16-18 m: mmrmmmmlmmgmmmmmmmmm\n8-11 t: tntmlpgfttv\n8-11 f: mffffffffff\n15-18 t: tftttttttjwtttjttjt\n5-12 d: pzdrdqxzqzwgzwcgsv\n4-9 d: cdwdmhksdzgd\n9-17 l: mlsllmgdlhlgmfglr\n1-5 r: rrcnr\n3-4 c: ccccccc\n1-2 g: cskvbgnlgnpfrh\n5-10 n: tfnnvnnqbbn\n1-15 g: fxpdggmtrqggdglh\n3-4 t: ftttppdtdbxzqw\n5-6 p: qpxsjt\n8-14 r: bkltjzprkhrhmrg\n1-11 f: gffkkmpflxlqfsf\n1-4 d: vddknn\n5-7 l: llllllvlllll\n2-11 b: bbbbbbbbbbbbb\n1-3 h: hhhh\n13-14 q: jqqqcckhdqkqqj\n4-14 p: mfwcbqxpdkhtdfj\n3-4 g: ggggg\n6-14 v: vvvvvvvvvvvvvvv\n16-18 v: vvvvvvvsvvvsvvvvvv\n9-13 l: llllllllllllllllllll\n3-4 r: zjrr\n6-8 x: xxxxxxxxx\n2-6 p: zlgpps\n3-8 t: ttjtbttjt\n18-19 k: kkkkkkkkkkkklkkkkkk\n8-13 q: hcvnkpnqfzqmqscwf\n15-16 p: ppppppppppppdpmr\n7-12 c: cqccccdcccwsccc\n8-12 d: ddhddddgdddgdddd\n3-4 x: xxxx\n6-8 w: wwwrlzcbwhwlqwwwv\n1-5 h: hhhhdwhhlhph\n4-5 f: fffvz\n3-8 c: ccqqcccc\n3-5 k: xkwkdkh\n2-3 s: ssssshhsss\n4-6 x: gxxxcx\n8-12 l: vsnlqbjltbtlphjtf\n3-6 n: tgpnqxnntkmntvl\n9-12 f: cjzkldhmvclk\n4-7 b: zgpdznb\n1-12 s: ssssnrssssscnssssq\n14-16 p: pppppppppppppwpjp\n4-14 b: xpbbvcbbbspgbbqbw\n5-6 s: smssrb\n5-16 b: tzbdqxgkbxbvzwvt\n2-14 f: vzplvvftrcdmgs\n13-15 t: tttttkthhtzttdt\n2-3 s: vpwss\n9-14 s: sssssssfssssss\n7-13 s: tsnwzxwzjzsjkchlg\n12-15 q: qlqbbfqqvqmcqhvq\n12-13 g: gggggggggggggggggg\n1-3 g: jjdqzgwgggggg\n5-6 n: nlnxnnnwv\n15-16 m: mmmmmmmmmmmmmmmjmmm\n2-6 m: vmrqrmz\n4-5 n: nvnnwb\n13-15 x: xgblxvxmpskqwvlx\n6-7 v: vvvvvvvv\n8-10 r: fsrrdprbrfxhrdv\n4-5 j: jsjjf\n4-6 k: kkkjkkk\n5-9 j: cflfvjjtjljjhbhnj\n9-11 t: xswktxtqrbm\n4-16 r: drrrrrrrrrrrrfwnr\n13-18 p: bbkplmplqmppppspmp\n9-15 r: kgnwvxmcrmltpcgwvrb\n4-5 c: jcwcccnc\n2-4 f: ffff\n8-12 d: rvsdddvdtdcd\n3-10 v: qvdzjhcvsz\n10-11 x: qcmvlvgxkxx\n6-9 s: sxsspsksjssfsnssrvs\n3-7 x: zbxxxxxwrpxgr\n9-17 j: hrppktjsqhzqwxhdz\n5-7 z: nzzzvzswz\n6-7 p: ppppsbpp\n10-13 f: fffffffffhffmfcxffpf\n9-12 f: ffffffffsffmf\n14-15 r: rrrrrrmrrrrmrrzr\n3-6 w: whwwwwwwwwwwm\n7-8 b: btqbbbbbbbz\n13-16 z: zznbbdmbzhjzzzzzmngd\n6-7 q: fqmqltdqqqdwqss\n3-4 h: hphhhhh\n1-3 r: rdgktrmrmdztrkmn\n17-18 w: wwslwwsmwswwfxwwwww\n6-8 p: ppprpppp\n8-10 z: zzzzzdzrztzzzz\n8-11 m: tmldmmwxmjmrmtrmm\n12-16 s: phpprsdvkmsssmssnqwm\n3-4 h: qplnhkhh\n14-19 w: wvxvqzqzdnjjwwtdbwc\n2-4 p: fpnp\n3-5 c: ccccpcccccccc\n3-12 g: ggwggbgfggrxggg\n4-8 r: rrszrgdbrw\n2-4 n: vnmgjtps\n8-10 j: mjjjxjjgjc\n2-6 m: vmmhmm\n3-4 w: vwww\n10-18 v: mvnbqbvcpkvnhqvvbvzv\n15-19 g: ggtggggggbggggxggglg\n7-13 m: mmmmmmjmmmmmmm\n5-11 d: ddxhdndtfdddndxzdqk\n3-11 f: ffnfbfffffgrff\n7-14 w: fflfpgklzlxfwwt\n15-17 d: zxllzgbjrrddnmdhdwp\n1-5 n: nnnnknnn\n15-16 j: ffkjjljgndljjnjscs\n2-3 v: nvvxvntqcrqq\n5-6 c: rcccgpcc\n2-3 j: qjjfrhwljgczj\n5-8 k: kkkkkkkkkk\n1-2 r: rrrrxrcrrr\n1-15 t: ltttttttttttttttt\n4-13 q: prlqzqvgpqslqqcqqjq\n3-12 q: gfbpbdrwrbqt\n2-5 q: qsxccx\n4-9 t: slglllcdkjzdmxt\n2-3 r: rlvrrr\n2-13 w: wwmkwwwjcwbcwzwwcg\n3-4 t: ndkvx\n6-13 b: qflzbbbcfzkbmtx\n6-10 j: pdckwfhkjxqljnjjw\n1-5 j: jjjjjqmj\n2-10 x: xlxxxxxxxfxx\n1-5 r: rvcnr\n4-10 c: cnchrcsgcvncpq\n3-5 j: fjxbj\n4-11 c: ccdlcsxckcx\n12-16 m: mmmmlmvmpcmnmmdbmms\n3-4 q: qqwlqqqq\n2-6 f: kvcfdv\n8-10 m: mmmmmmmmmm\n7-8 j: jjjjjjjj\n5-7 c: ccccncsjc\n7-9 f: fffffrpfplhzfff\n1-6 c: sbcccscccccccccccc\n5-7 j: njpjqxjrljnqjwjgchtj\n11-18 b: gktmxhjwbmbpjrbqkb\n1-4 f: lffbfffffw\n5-6 c: ccccsc\n10-14 v: mkbxvxvgvvvvbv\n5-6 x: jxxxxxx\n4-6 c: qccmhcccw\n3-7 s: qwmtwprsxd\n10-19 d: rkdwbrndmxdsknddddl\n2-5 p: pppppppppppp\n6-7 w: wwffxwc\n2-7 z: pkzcnxr\n5-10 b: bbqbkbtdbbbbb\n10-13 j: qmqgvhzkjmfjhjjm\n5-6 n: nnnnrn\n7-14 n: qgmlsvnjnrdzcdbstxv\n9-16 q: qpllvqpmqlqvkhwqqrsm\n2-4 s: sqpnqsgsssq\n13-15 z: zzzzmzzzzzzzzzzzzg\n10-15 b: bbbdbmqbkjbbtbd\n2-5 h: qrhhhhp\n11-12 x: xxcxhmxxxxxx\n13-18 c: ccccccccccccdccccccc\n12-13 h: knktllrdtgjhhhxpd\n9-13 g: tdpfvzcqgbslgnbxkgpq\n8-15 g: gggggggsggggggxggggg\n1-3 v: wvfvvdvvvvv\n10-11 b: bbbbbbbbbhzb\n15-17 b: bbbbbbbbbbbbbbpttv\n1-7 m: bmmmmmdm\n2-3 h: rhksbfgh\n6-7 w: wwwwwgf\n3-12 h: hkhzhhhhhhhlhhmh\n3-5 x: xxcxpx\n7-9 m: mmdmtmjdxmm\n1-4 n: lnnlnnnnn\n4-9 k: kkkmkkkkbk\n3-4 t: tttt\n5-12 k: tkgwkrwwzcwkhkmk\n4-5 f: fhfvfgwfff\n3-5 h: fmbgjrxgh\n5-6 j: xtwnjj\n11-15 r: rrrrrrrrprqrrrfrr\n13-14 f: fdffffflfffvfff\n1-6 b: krsbnd\n6-7 b: bbbbvxqbl\n1-2 v: vvcvvvv\n7-9 t: ctfcmttht\n2-7 v: bcxhtdwbfqvv\n14-15 n: nnsnlnvkqztnxnn\n2-6 g: rknndvrhbg\n2-6 z: mzzzdzzzz\n11-14 q: qqqqqqvrqqzlqkqqq\n3-4 t: ttcrttjtknjpvwhzm\n8-10 j: jjjjjjjjjjjj\n2-7 c: ccnhcccc\n3-14 m: mmmmmmmmmmmmzmm\n10-16 w: swsmwwwwwrnwmwwk\n4-5 m: bmdjm\n6-11 l: cmlllxllllkrllf\n14-18 t: tttttttttttttvtltt\n11-13 f: ffffffffffhfzfff\n3-6 g: wkzxtrkgr\n7-10 c: cccpcccccvccc\n3-13 m: mmbmmzbcmlwmmfnmmq\n2-3 p: pdszp\n13-14 r: rqrrfrdrrrwvrrrrrwrr\n9-11 g: gggwggggfcpg\n3-9 p: pnppppdpppp\n17-19 k: gzkktkxkkskkkmkkkpk\n4-11 x: bmqhtxzqmvx\n11-14 v: cvsvssqghzjqxm\n1-14 z: zwxczzszzdmkztl\n8-9 t: jkpptwctt\n3-8 z: vqmzzpjj\n6-10 s: msdrxssssswssjz\n1-7 z: zhpzzzz\n13-17 m: mmmmmmmmmlmmgmmmgm\n9-10 n: qgtcnnnnnkcnnn\n2-4 v: vrcvv\n4-5 h: hhhdwh\n8-9 c: zcccccccc\n13-14 p: pjxcpmdpppmgwt\n2-6 c: gskmcc\n6-7 q: qqqqqqqq\n3-13 b: gfbxqtnxnqxbb\n8-10 l: lhrlqlrqlsllll\n5-8 z: jlzgzzzq\n10-13 z: zzzzzzzpzkzzgz\n6-7 d: wddnddd\n10-11 h: vbhhbhhhhbz\n7-10 s: sxssksskqs\n1-12 b: vdgxvsxpgbbpd\n5-6 n: npnvckn\n3-4 v: vvjvv\n9-15 r: njhhqlmcmncjrpx\n2-8 h: bghrdhhrjbhhhbjz\n5-7 q: qqqqlflrq\n9-12 f: qfhfffffffffffff\n2-3 m: mmmblxmbpz\n13-16 v: vvvvvvvqvvvvlvvw\n6-12 b: wbfgbghzbmvf\n12-15 r: rbrrrrthrrrrrrrr\n4-8 k: pkprqkkkpsphwnhp\n8-9 h: hhhhhhhsl\n1-3 b: bsbn\n15-16 j: ktjjjxjfjjrjjcjk\n4-8 z: smzzmzzszzr\n3-4 m: qmmmzqzlrpt\n11-12 n: nnnnngnnnnwgnnnn\n5-10 x: slpbxqfxxxx\n12-19 q: qqqqgqqqsqqqqqqqsqqq\n4-10 s: ssstssssss\n6-11 g: gggggvggggzggggggg\n12-14 c: cckccccdccxnngccc\n3-5 j: jnrjvj\n2-6 t: vtdtms\n13-15 p: ppppppppppppjpq\n1-9 j: jjjjjjjjjjj\n8-9 d: wfdddddkddxd\n3-7 p: pwpcpvp\n3-4 h: hhzl\n9-10 n: lnkjxnsbnn\n7-11 z: zzzzfznzbzbzbzd\n1-2 m: mmzlmkm\n6-7 t: jbmpztx\n5-6 v: vvbvvtn\n7-17 q: lvdfpbqxbxdnwqzjlfs\n3-5 z: zdzzzxczbzzz\n10-14 f: fpffffffffffff\n3-8 x: wxxchxjx\n1-6 x: xxxxxx\n14-15 x: zmxtcbjxtvxxxxg\n4-5 f: qfffghf\n2-3 p: pppp\n10-12 n: nnngnxnnnhnsbntln\n1-2 w: ftwwww\n10-11 f: fffffffffpkf\n4-5 n: nnnnnn\n10-11 r: rrrrtrrlrrrr\n5-6 h: hjhxhhnbhlbxvhs\n1-4 q: qqqq\n7-11 p: dkpfppwzpxqqx\n11-17 l: lzknllnllkgvllllh\n1-3 z: zzzzz\n16-17 b: bbbbbbbbbbbbbbbbbb\n2-4 n: qcpwnhn\n1-4 v: svvhvvv\n10-12 l: llllllllllljl\n1-2 l: xzlllllll\n1-3 l: plklll\n4-8 x: xxxxfxbxcxxqxfxpx\n3-8 b: xgbbwncxclbv\n5-6 m: mmmmmmm\n4-5 p: pptppfnpl\n14-16 x: gwjfkzwrldzxpdxgk\n10-11 z: zzzzvzzzzzzz\n3-12 k: kkxkblkkkkkgkkkmkk\n4-6 p: pppfpx\n2-7 r: rrmrrrrrrrrrr\n7-9 p: pbpfpppppp\n2-3 v: tvvbbjnvv\n5-8 j: jjjjjjjjj\n2-5 c: xcwcczjkvgccccvpkd\n5-7 f: ffffpfqfffff\n11-13 m: qmmmmmmmlwkmmmmmm\n14-15 d: dddddddpmddddszdd\n6-15 b: kptphpbxxqsbxlhxplml\n16-18 x: xxkzfxgqxdcvzdxlzn\n10-12 m: clrmmmmksmmm\n1-3 w: wwpww\n8-15 g: ggggggnlgggggggggng\n3-6 h: tnhhxh\n4-5 l: lllrdllxlnllqll\n1-2 m: lmjhvl\n1-4 k: dkkhkkk\n6-7 x: wncxcxf\n4-5 w: wwwhfww\n4-8 b: bshnbbdbmfxbbrc\n4-7 t: tbdtztm\n7-8 n: zntlhccn\n10-14 f: frgffffffflfzlf\n3-4 b: nzbbbnfbdg\n2-6 m: srmwsmtp\n5-6 z: zzpvsrz\n2-13 m: msmmmhmmmmmmtm\n3-4 c: ccvcc\n1-4 n: nnnnnn\n2-3 j: zdvsj\n4-6 c: ccctccc\n4-5 z: zzpdztnrzl\n4-6 z: hzdzpz\n2-6 x: gxjxcxxjqcb\n4-9 j: jjjvjcgzl\n5-8 z: dmppvjqfzzzzn\n5-7 v: vvvvvvr\n2-5 l: lvkkft\n2-11 n: nnhnqnnvnsn\n1-3 d: dhndddjfd\n8-9 x: xxxxxxxxxxxxxxxxxxxx\n4-9 p: fntpgcwtwv\n4-12 q: qbrlqzqcnwtjq\n2-3 f: ftcjvnfkzx\n1-2 d: dddd\n7-10 g: gggggghggzgg\n4-6 g: zhgbgphggtm\n2-4 l: lblq\n14-17 n: bnnngnnnnnnznnnnn\n5-6 f: fxxffffhmswq\n6-10 p: pppppppppp\n2-4 q: lsxgqkqsblqqq\n11-19 x: xxnrgxxxxxxxxxfwmxx\n4-7 j: jjjtjjljjjjjjj\n6-7 l: llllllll\n1-4 f: pfftffffnffff\n11-13 r: xrrlrrrrrqrrzrr\n1-3 p: pvpqp\n9-13 d: ddddxrnddddfdscsd\n1-2 f: lbbr\n3-7 t: mpttsgwb\n2-10 c: jccchctshcpchwx\n2-4 p: lmpd\n1-2 n: zlnxnnf\n2-3 w: fwwtqw\n9-16 s: sssssssssssssssssssj\n1-5 s: hssrm\n15-17 b: ghpvwjbzpksvhsjbj\n8-9 h: hzrvxhhdz\n3-4 j: mkjjpvkfhprr\n1-8 n: cnnnnnnmnn\n1-4 x: lxxxx\n3-5 m: pjttcvvcmfvzffcfmmv\n3-10 r: rsxtvwjrfrqrrbzzr\n2-7 h: hhckhhr\n7-8 r: srrrmprrrhkzsndrrkr\n4-13 m: pdsmlsnxcmhmmsvc\n3-12 d: ddsdddddddddddddd\n14-16 s: sssssssssssssjsds\n3-4 m: djfp\n2-5 w: bgwpc\n4-5 r: zffrrqrs\n4-9 r: rrrgrrrrqrrrrrrr\n4-6 w: wxwwvw\n7-8 d: tvddddwtdddldd\n6-7 g: tgdggggfzrgqggggmggg\n2-6 t: tgtfdtttttttttqxd\n14-18 f: fffsffgfffffskfkff\n1-12 b: qbvbbbbmbsqbcbgbbbcp\n5-7 d: ddbqhdddsdm\n6-9 n: tqwnckjsn\n1-3 n: nnnn\n5-10 w: wwwwwwwlwwqwwwp\n3-4 r: rrrr\n2-12 m: mmchtqmhfdpm\n6-9 d: dzgzdbddwgdwldb\n10-14 k: kkkkkkkkkkkkkzkkkkk\n13-14 v: vvvvvvvvvvvvvvvv\n2-12 f: btqtbrvkmfzhz\n1-5 b: bbbbbb\n7-10 j: jhjnjjwnjbj\n2-16 h: kjszvtwrjgrvzrqlzcb\n6-8 t: nxzftvtnrtxg\n3-6 h: cvhhnhnhhcgst\n5-8 f: fxgfjnjzz\n11-12 b: bbbbbnbbbzbqbbbq\n1-6 h: hjzhhw\n18-19 m: xrvmftmpkpmrlkqmxvm\n6-7 v: cvxbdsr\n5-6 v: vjvvvvf\n2-3 m: mmzpm\n3-4 k: mkghkxc\n1-6 s: sjsmjsstsss\n2-10 s: wndvkxbmffh\n9-15 q: qqqqqqqqqqqqqqqqq\n4-5 f: rffffpc\n10-11 v: vvvvvvrvvlmv\n6-7 n: tnnnnlq\n16-18 g: gggggdgdgggggggggggg\n5-11 l: lwlfvzlsjll\n11-14 v: vvvvvvvrvwvvvv\n12-16 j: jjjjjjjjjjjpjjjx\n4-6 w: wwwwwr\n2-4 d: ncdd\n15-17 c: nccccctccwccccqcqcc\n3-5 m: mnlmqtsvn\n7-8 w: swwwwwnl\n6-12 l: zlplklqldllxlpvb\n7-9 h: fhkhqnfhhzrhqhhh\n11-14 v: vvvvvvvvvvvvvv\n16-18 v: vnvvsllmvswcvzqvvxjl\n7-9 k: kkkzkkkkskkkzkkkjlk\n5-17 w: wckwxwlrlwrncxwwb\n2-4 v: vqvvqgx\n2-3 b: bbbbbbmbbbbqb\n4-5 p: ptphbprpv\n3-6 q: kqqkpqbnt\n15-20 n: nnnnnnngnnnnnnnnnnnw\n10-12 p: pppppxpppmpnjp\n5-10 x: bhwlxkrdxxqvkphvmfgn\n2-3 l: lllhql\n4-11 l: wllvwfccllk\n7-11 d: dmvmmzlngpw\n2-6 v: rskbvv\n7-9 k: kwtckkkkx\n11-12 j: jnjjfzjjjjjjjj\n2-5 v: rvvnnjjxjnnt\n2-6 j: ljcqpjmhgmmcxjkgd\n1-3 k: lklbx\n14-15 c: ccccccccccccckc\n2-6 x: xxxxxxxxxxxx\n1-7 t: btrtxtjgtt\n2-3 t: sttt\n5-7 k: kkkkzksxgb\n3-4 q: mnnxqwtnw\n8-10 t: vwzmnxjdzx\n4-10 f: svglffffpzw\n1-5 s: dsfjsv\n8-10 m: vmwpmmmjmv\n6-7 h: hhhhhwxhhhhhhhhhhhhh\n7-19 n: ktlblznjctnnrzwrknt\n3-6 n: nzxclnvc\n2-5 t: wtrttkx\n9-10 j: jjjjjjxjjj\n12-13 r: rrrrrjrrrrrwx\n11-16 c: plvcnjcpnkvcpcmcpbc\n2-3 c: gccclcslqdmsg\n10-13 d: dddddddddhddd\n8-9 w: kklwmwqmwwtwqw\n2-5 q: qsqqwx\n3-4 c: hcqvfbccccctcc\n6-7 k: kkkwqgv\n13-15 h: qhhmhkjvhhhhngnhhhhs\n4-5 d: msjdd\n6-7 g: ggggggg\n5-8 r: cqrdrfxnrrjrsr\n9-15 r: lprrqrrrrrrrrrrbr\n6-7 h: hhrqhrlhhh\n2-3 g: glgg\n1-11 n: nnnnnnnnlnnjzwnnnn\n3-6 l: bwllrl\n2-6 p: nptphqv\n3-8 l: lrlknvlpcm\n2-8 g: wjbgpjrlfsgg\n5-6 w: swwfww\n1-8 p: pptpxppppppwpppfx\n3-4 d: wlddg\n15-17 c: sccccccgccccccccc\n4-6 n: nnnhnp\n8-9 r: prvrxmprlzxs\n9-12 t: ttttttttttttkrtt\n5-6 h: hhhhkfh\n1-10 q: psqpzpkqtcq\n4-7 f: ffflnfzkf\n10-11 j: jjjvjjxfjjj\n4-11 p: wlhgrxlclnt\n1-14 t: wfntvttxmldhqg\n14-19 c: ccpczqrccfhtcvljlvc\n11-12 w: wwwwwwwwwwwww\n2-8 z: lzsdqxzz\n17-18 p: pppppppppppqphppprpp\n2-4 j: wjjq\n4-6 f: ffkfffgff\n8-20 x: phxxxxxnwxxdqzxxxxxx\n9-10 d: xvddddsldd\n3-10 l: rglllxlnfl\n6-15 d: ddddddgddpddddwd\n6-7 w: kwwwwhs\n8-10 r: rrrfmrrrrd\n2-11 x: xcjcgtkzkhtmrxqjxxx\n4-5 v: wvmdr\n7-8 d: dddddddd\n7-13 l: lllllldlxllln\n13-14 z: zzdrzzzffzzzzzznzz\n3-4 j: zjjj\n8-9 s: sssssssbss\n6-15 f: pfsfvfxqpffqfqf\n9-11 p: pjpppprzpcphvppwp\n3-4 c: ccrd\n9-10 c: qcwbccccjhwvtwcccnc\n3-4 x: xgxx\n5-6 x: xgxzxxxxxxx\n5-10 l: kxdgnqrkpqgcmcmnk\n4-7 v: vvvpvvvv\n1-4 n: nqnnnnnn\n5-6 p: pprgpst\n14-20 q: wlxrjczhxwdctvpcgxqc\n5-8 g: gcggbggggggzgg\n10-11 g: ggggtggggggg\n14-18 h: hhchhhmhhbwhhlhsvrh\n3-6 f: ffmfffffffff\n1-12 n: rnnnnjwnnnnv\n7-12 l: bpxlhthjplljwxvvvjm\n4-5 j: jhjjjjj\n9-12 q: qvqqqqqvqsqq\n3-4 d: ddnvdddddgddd\n1-7 h: xhhhhhkh\n1-4 f: fffff\n3-4 q: qfrq\n5-8 k: kkkkgkkqkk\n14-18 k: nkkkkkkfkkwkkkkkqk\n1-4 q: qqdc\n17-18 f: ffffffffhffffffffff\n16-17 b: bbbbbbbbbbbbbbbdm\n17-18 z: zzzzzzzzzznzzzvzmwzz\n3-4 g: tggd\n4-6 t: mlttlqqz\n11-18 f: fgffgwlffffftfzfrt\n7-9 h: hmmhhhpljjn\n16-17 m: mmrsnmmzfhtmrjwmmjm\n2-3 l: cllcbm\n3-4 k: fkkkqbpbmmcd\n9-10 f: kfmchwkwzfh\n5-6 j: jjmjjjjjj\n3-5 v: vsvbfvvh\n3-4 k: tgkk\n2-8 k: lscvjjnhpxl\n7-8 v: vvlvvfkqqv\n8-9 q: qcqqtqqhqvqqtq\n12-14 z: zzzhzwzzqzzzzwzzlz\n5-8 k: kzknkkkkkp\n5-6 s: spsssw\n2-3 k: kmkh\n9-10 f: vnlxmstqbsg\n1-2 k: kkkk\n6-15 r: rcrrrsrrrrrrrrzrrqrr\n3-8 c: cfczjwkcxbc\n15-16 k: kkkkgkkkkkkkkkkh\n5-18 c: ccdcxtccccccpcfccv\n6-11 w: qdwnfhzlwwwwxwwgwhwk\n1-6 d: fddddddhddw\n7-10 p: wpbppvpzwc\n12-14 k: kkkkkkkkkkhhkslsk\n4-7 r: hdrrfhrqrdmcbnlrrjbr\n7-10 b: bbbbbjbbbbbb\n4-7 p: pzppzppsp\n11-12 h: qhhhhhhqhrlwlhhthh\n6-9 s: sbssxsskswvp\n17-18 b: bbmpbbhhbnsbbbbbbb\n8-9 g: wgggktggg\n1-12 s: vsssssfssssrqk\n2-5 s: bsgssxc\n3-4 c: ccccccc\n2-3 d: dddd\n5-7 r: ldlcjrrlrngr\n5-10 m: msmjqmmmmm\n3-7 c: ccccccccc\n13-14 p: pppxpppppppppcpppp\n2-3 r: jrfrvrk\n11-12 g: gggggggggggg\n6-18 t: ctvqgcrgnxdvbzjfrrbt\n6-11 f: rnjptfnwgxfp\n5-8 w: wwwtwwwwww\n14-15 r: rkrbrvrrrgrczrz\n12-13 t: trtjtttlnxnxx\n5-8 l: qnwllfsl\n2-15 g: xgtcjftlqqfwkggpf\n11-16 j: jjjjjjjjjljjjjjjj\n8-9 b: bbzbkbbvgcbb\n5-6 r: dvkxrrsvrrksszsdr\n12-13 j: jjjjjjjjjjjhdjjj\n16-17 z: mzzzrxfzzzzczzgzz\n2-15 p: lpjxcdzjmnghfppr\n9-15 s: ssssssssnsssssss\n1-11 t: tfvtqvlbtld\n4-5 k: kkkczkkkvkkk\n2-7 p: ptphppvppppp"

	utils.Log(true, "info", "Solving test input part 1")
	day2part1(testInput)
	utils.Log(true, "info", "Solving real input part 1")
	day2part1(realInput)
	utils.Log(true, "info", "Solving test input part 2")
	day2part2(testInput)
	utils.Log(true, "info", "Solving real input part 2")
	day2part2(realInput)
}
