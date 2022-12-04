package main

import (
	"fmt"
	"github.com/adrien3d/adventofcode/utils"
	"strings"
)

func solve(input string, v bool) (part1TotalScore, part2TotalScore int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		itemPriority := int32(0)
		half1 := line[0:(len(line) / 2)]
		half2 := line[(len(line) / 2):len(line)]
		for _, itemP1 := range half1 {
			for _, itemP2 := range half2 {
				if itemPriority == 0 && itemP1 == itemP2 {
					if itemP1 >= 65 && itemP1 <= 90 {
						itemPriority = itemP1 - 38
					} else if itemP1 >= 97 && itemP1 <= 122 {
						itemPriority = itemP1 - 96
					}
					fmt.Println(line, itemP1, itemPriority)
					continue
				}
			}
		}
		part1TotalScore += int(itemPriority)
	}

	for i := 0; i < len(lines); i += 3 {
		itemPriority := int32(0)
		for _, item1 := range lines[i] {
			for _, item2 := range lines[i+1] {
				for _, item3 := range lines[i+2] {
					if itemPriority == 0 && item1 == item2 && item2 == item3 {
						if item1 >= 65 && item1 <= 90 {
							itemPriority = item1 - 38
						} else if item1 >= 97 && item1 <= 122 {
							itemPriority = item1 - 96
						}
						fmt.Println(lines[i], item1, itemPriority)
						break
					}
				}
				if itemPriority != 0 {
					break
				}
			}
			if itemPriority != 0 {
				break
			}
		}
		part2TotalScore += int(itemPriority)
	}

	return part1TotalScore, part2TotalScore
}

func main() {
	testInput := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"
	realInput := "GwrhJPDJCZFRcwfZWV\nLjnQlqNpjjmpmQlLlqNfZRvQcTWcTSTTZcSQcZ\nnNqjdspspngnmjmslqmjjjCDGrHPHMGddGCMCGPPPJWC\nGwmVZmPWWFFmBbVbZVwmbPsTCnlgQgnQfhlffffZnlQh\nDqVDSqqSMzLLDDNSHHLdqSdSllCQjsTlClhlflnTlhjgfgfM\nVHJztNLHGtcbvvPG\nbjrPrNCtNrjdcCPpptfpTVspDtfTtB\nJGQJMJQMmmmZMnnLpLBTpHCD\nWJJqWRgWlCJZhZRCQZwdPScdrPNbvzPzwvqz\nQNSQNBWQNLjZBNNhLhSNRsTcsrTTVzcwZZZsfrrbwb\ntCFtHpppppMldpvpqnMFmMVGrbPcrwbzswrzcccfvTfw\npdmCpgqCdmHHdJVWgSRNJDRVVj\nsNrFnvNSzrjQtQjQTj\nlcPmcJDLdPDbJPVLljdGGBBThBQTGwTtBw\nPDLqmJmpJQfFqfqsCM\nBnhctqdnqnRcBnslCJJCMrJrsG\nwNDMZpbQwMpCvCGVjlss\nWfzNwZFbwZzZmFZbLzNwzzzzcdqgRMTTPdHPTTPMRdcWgRPt\ngrsrVSFSSdFSDFVFjZZWwpWpZWZplgZZ\nmcBPPPBLBfNdLlvvWljWJC\ndMcmcRdbRzdVhFthSsTShM\nbzvJZMTzTZSHLCCdDzmDcc\nhqBqWPFssvshWvvssNqtsHftmfpHfdcdDGHmcpfctL\nWvBQgNNNhghTJbJQlJTZlT\nchcdwNwdbCbQctCjnnQGHsQspMHMjG\nnSSSJqJZzJgWWRfZDJSnqvTTsVvvHVPpHVfpjHMTjP\nBZRDRmmrDWSrZWWzWSRNhdnCFwChclFtwbNdtr\nlNgmssCtqLwqCCtfsCLHPFhhhmMhVzBDbVzMDMVz\nZnRlQTlJzFQFQFVV\nnpZJvRRGZSnWvSvrSLglsClfpfcLgNgpHf\ntVtqcVqFVtZhcfFtqlgSpmpDSDNPzSzZmNpw\nLRGTHqbrHdnGHrTCSSwNDzMDwPMzNwbp\nTqWGJrGHCHnTWnhsWcFthFjtfQch\nqNnTbwtctvffFcqfrHjMrFjVHRjSjZDZ\ndLLzWWPmCmCzGdsLgBLGGBDRMVMHRlrrrZDDZsNMrNNS\nPJQWggCzWNWJzGWfchvfTbJvfnnwtf\nghzdgzzdQsdqzzhMNqQzvhgQnRRBWTjWWGTRGWwGTZhwGnBT\nfsrfJHbFfDFLVLVFHrWCWrBRZZTGCCjwWZ\nHLLllcDPbLPQdPspMNgvMt\nfNDJqdPNbtHpCbwpCCCp\nRTMRLrzGrMRMRPWnnvSmgCHFCCFmmT\nWQsWQjzGWMsGQzWclQtVBJfBftNdtqVPfP\ngbTCVVmDVFdsgmgrrcfwlwfTfPlcRR\nqhQZqQvnQhLQhJnvfPcSwSwlfjGcqjqj\ntLNZLZZJJZthpzhMZDCdFmFsmWWmtDDgsw\nbqCvLvLppzPzPPvPbFztFtttBNGdGsRggSgGSHDdggHSzNgH\nrMQpWfMfrcTjWJhwWHHsSBsRBdSTTNBgSR\nVwfmWjwMWwccrWcWpQQFnFtlCqmltFnFLbbZmn\ncWqsMWJMzqJJMHsJcqsJqTqjSbLBdfdSbtzLbbLfbSfShfhd\ngplGvQmRrCrgZSZtSGZZjhbj\nCQmmmmNQRPvjgRClCvmmcVHPqMFMFsWJVqFFcnTJ\nQHHqvGwjjWNqvGTQGvTFcGwJRJbszcPtDbJVbtPzVbDptp\nMLdrgmSgZZdhdfbLVRpszlRDstRL\ngdSgMCSfdMnrghCWGRQvHwvNHjnjvv\nRDBZwvZBrMlsvnlb\nWdFQqdjWWcHHPrwSPnnSWnSS\nmLdqgqHmcjHHjqLHjLppmhfBfgtDtBJZJfVtBZwGZB\nCCWRJQnZlHtHtNZRFDcBhrcvhDrJVVDv\ndPPSqLzfsqGLSTzfLzLGdLMVVgvBcmgMVwmmDFrVgmBBBr\nSFjdTGzqpjdRbNRNnjtnQR\nhjNcwBDDwDFcjdFfjtFhtcRsGGgTsGRRRTsGGqZGRq\ngbmrLnbzLmvQJnQVVpqZTqzWSCRpqRTsSR\nMQMvVMbPQQHrQMnMPldtwNNfgHtlwBhdwj\nzwzwpzMfzrBMWfCCZrwzrMJDGGGnNmGNZvgNZsDDsGsG\nFbFqSbcSbSHqTjmgGFnJglllsDJm\nTbhVdVjqdtqTjVHqjPdthPBBWpCnRfwRPRCfBCCnWR\nhlpmbfJJpCSChmJMmrSjTjcSdjTtQQTtTtjF\ngqrgsqLzgnBgZGzHBnnsQNNQtjjcNNjjtNFQNcNH\nLVRzgGGzzzPCVrJMbPJb\nVHrmqFnVdvlzzNrr\nPMtwBJPBcPwfbwBJndplLvLdLlgMMzLL\nbBZnTwbtnScfQJPJwPTjqGZFsVFjDHHGhhHhVj\ncftqScHJrfVfrrRZ\nDTTsDvvlBbTGrWBwwsWDBbWdVpZjjZjpVPPGhRRVjVZNRPNN\nlsWdWDbrTLBsbdrmdwbMJtmHMQJccFHFnJFqFt\nSWNPTPVSWChCSmQQhpppJdFJLpDpgLJmLd\nNGGtNtGfHtDpdJdqLB\nNcsNGNjHZsZGnzZfnGhQnhPClrVlQPhTVVhl\nQDdgMBsNhhMgcWbZdzmWLzFzWH\nfRqRJJqGCvrJGjCRRrSJlfPtHzzPmfFbtPtLZZLnmt\nVjvwwjlwVGGqJSSqJFccshpgNhQNQTsVgBgT\nwvDLDwCbFgSTfTSJJgfB\nqsRhmhqchmVhPdfTHJSzpCtJpfPf\nhmdhrWrddmhlqCRcwQjDLMQnMFDZnlLl\ntrMWtlwwMplMZMCZWltDpzBLBnflVLBbHzbBSGlVlL\nghhqJTfmjQjfqqznznnHnBRzBLmn\nsQhPQsjjQcQcTsPqZWwwZcFfWrWcrZww\nMRVpVCZZTHWVMCHvgNvVvbQSqgQSlg\nNFmnrNDDfnjFnndfssmcStvjvQQlvzvllqvwQllj\nGGPNmBrFNdcfcGrsGcdmDFhJHMMhHLZJMhpLHCMMMMPJ\nDSvDGdGFlGGnDZFdVSZvfPqwnfhpnrqpPNpLPrrh\nsWcTjtHCsTmsCNfgMPjpfPhqhP\nBtHzBzChzBBvFSDJvVzFJJ\nsfsNrsFFBTfjwwtNNWHPVCVWtSCDDCDmmS\nzMdhMMZnSccMmmWVWmCPlC\ncLSScJZQbcvLhZvnzBwfTjrpNwNrBFffpb\nTBrCBgrTngVQBVbhrCtgJJrGssGsMGRGcjMcNjfN\nLZdSLvHMFdzFRWsLjcGRWWNJ\npHpzlqPqFPvdBthgMbVPDhgh\nSZlnZZvBvvMrcBnllBMZSvhGMtQwFMGztthfwQtMwwPf\nHLqsDgNsDLDDDjggHDHszthzFbQGTghPGQPbTfFT\ndmLqDqCmFNjJsjHdssFNHDVWZccnRllnVZvRSBZrZlCc\nSccnnSGGftShfHSHHhnvbMjvVlCjzbVzzbMMTbCB\ngRpppNNQLWqZgPZwNWwwBMBbDlZCTzVTjHMMbBjV\ndqNQPQRqrqpPcGtchhdfhHSF\nmfDzgnNMMszBtJCpHlrjnFppCdHj\nLLRThGGZcbClBQpdWFGl\nbSqVTbBbMVMsNmNM\nBTTbbLVpfchmjbsj\nJSQJHDMHqdNZTZlhFFhCFFrNhNcsrr\ntMwJQlwMMlQwDDJtWGLGPpWLLGnTPn\nLcVQQCPPLqTzqQTcllTzhnHHfFJRcGHcFfwRGHwJjJ\nstdWDDBtVgbpWgZbsNgDNdWFGMnnwHfjHFpfwwMGMMGRjJ\nZWSDtgNdWNBdgsdsNDDsdbDlTzCVSTCqQmSqTQSvhqLVQq\ndZbgdZbNtmqttFJtHHzcczMcFszHnsvH\nwwpQplQQwqVVjqwPjCGCSMCMcHSHvvzHMzvcsrMc\npfjlQRpPRRLQWtmLNdWdmqqJ\nCPTPPmbjmVjVGCvzbjjPrGsnnMpttdtGdncdMccDRd\nlhlHzQSHwzhJLwgWgpMDMMsDdcDQMDMMns\nBHZghLWwSFBJJBFvzmbfjNZvZmCvmb\nPBGcvvcRwpwNcZcNPpPNcTHGdMtrCWrCCtCLWMtWgbVdMV\nfmsJjnqmmfsjQJnjFzSFSqsqgWrtMttZgMWVMbbVMdbSrLtr\nqQjjZFmfjZhZmwcvPhNpTNBTwN\nHHlVVmmsbbqMsJmVzGSBMSrQQrRrGvvnDn\nPZcphZPPZPhjcpdWgPZhRPfcDSrtDBSGNvtggrQtnvQNGNDn\ndcWwFjpcPhRcCpjwdCPLzHblJbLbzmsmbTwzqH\nhRfzTTfRrTGzhGWTrRrbfcQZQSttWtwddJtvdJJvWSHq\nnpjnDjFlpDnFFNMjljCnFMQtHHtqNHNQJwwZZqstNwJJ\nDCjpLjjpVLDMDpVLDLQbbhzBhVrcVgVGQQcz\nLncLBLjCSNrNrNpCLQBBBGwqQwzlzmggvqRqgllmzwtv\nfMZPHhhHfthMdbRgHJzmVqlvwlwg\nhfsPbZFPPDsfGLcBtSFNBSjL\nMlZmszBMJBHrMBMbShwSFpbZSZfwwb\nTCLCcPNGTgTPNGWtCtcWtPcSsRfRjRwjFbfpNFDjwsFspw\nnVtqqsWsdHzJHqmM\nRCrhSmWrmrvmrvhMvRNrRCzCJcQQbPtsMZVGJJtsZssPcQcZ\njLFBGqLFpqBLgZVbPbsLJQcbsV\nHjDljGFwrRHRRTrS\nGZZhnrwZBwNjRPRCbCbn\nfJtJJpsVfpgNTbVNFTRP\nJJcpLJfLdcWLdplwRdQMBvSqwRhvrG\nwmZDPlRlCDwglgsHtsBvdBHLFLSddr\nVbVMnMftfVjQWFFHdMBdBFMFHr\nzfjtnGqqnjGqfjPcDPlZPlRDzccw\nBRjhfhvRgnTMlFDDJfZzZFFQDZ\nqLdqcNttwwcwwSPSpqLNmrwmrZsGzzDFZGZFzVssrzJGnsQG\nwSNdHScScdmwHSpdNcmmtLMvChRHbvBMTBnCBBvhvlCh\nJgWTPfFPgCPPlCntQSGghHvQnSdQ\nBzvMZvLVQpdQpSZh\nRwVVjRDVcRDNDTlJPqTv\nSGHSrBBRPhPPHQcTccQTRRQjTN\nvvWvspCbzWVWVrWdjj\nwZpDzCDgDbCZJZzJGlrlqPqnqPllmH\nFCncCrDWMLCbjMCcFpLdzZfmZzwwWzdzNRZdWB\nsqsgTqHSqllNldMwlZzJ\nMtHPTgQhvhhqcrDrrDpjLCQc\npPPvmPWSClqqPvqCmSwqmgGBWDjhGLHfjhDLJGjBBhNj\nzrbdcdMndcRdTrsMcbTRdzRFVHjLjDjNLNHsfDhNGjhJNhDj\nRdFFcnTdZcTrRRdFFbZtwQCPQglvPlwJwQPZSqqP\nwlmbvwmvQvWQsvmbsSsQbswlRCNPfCTcTRVCffPtTSCPNRVP\nFhJJJFgFqJGBtDpJhTTcVcVhdcCdCdTV\nGDFtgLFnqqDGqGZsQvsllrjbLjbrvw\nlnFSnJvmgvLlfnJpgnsjnjgfDQWqCJqZdDtDCtCtCdDrtDDQ\nVTBBMPFcNNtMZDMW\nVTGbzGGhTbTGHwVPvvFnfpvjgHnfjppp\nJJwHqvlvDjljDwJFlZjZDwHNNsMqhNpphNpmNVzpsnsnRV\nmTLgrLLcLSTTTdmPPfrrrnssNhRNWhgngzMWzgzVnM\nSmTfdSBbBJbtjJvljl\nbPNLwTCLLQQqtJsf\nzdnnZVlWWGGRWGWdgdSStQMqJSMRptftbsMf\nFWbvgvZZZZgnTmwrrhrFPCrP\nHcGzzszFGllHWHbZspHbHGsHTwwrTrLLCNjSZwNjNjjCCNLj\nPBJMJQJDDDnDggRhMdRSLmjTmTwwVjVQSvvwvC\nRqfdhgDPDJDqJJnBdfzWWHcstslcbtStfHzl\nzvRRlCqrdNdZcZpjBpVwjsmjsm\nfgbTDqbhGfDnLDnLLqLhFmsHpTPHjHppppBwpwws\nnhnnnDDngDtDbfSbDnGhhgRlNvQdQqNvQvtcQQNJRNJN\ncZbCcbbScCbcmPGjPfSBQQSq\nlnMnnVsMVvmzzGMDzPDf\nLhrTsTTglrnsrrWWVvlwTnNtcpZRCmhtbCZFdttZbRCp\nNWrFPZVWNVrvvrhtnNdddtpldmjm\nDcBQBDsJbCwQnbtdzmjjjljbpjbz\nqCDcGsDJGCcBDBcswJnBJQDfWfqgvZSvgZPfrVSWvPvZZZ\nvcsdHdGtHtMHMFtVsddsWCcbppZwjScLpWhbjRWR\nNTwrnzJrgTPrDwnlphRpjSpWbJJLLZWj\nTlDPfPnzzlzTBzzvQFFBHMtVtqBqqw\nNHnqqfZvZBNHHvgfrSlJrJCSllJRVrCn\nTDTdhLMWjFcddMJPSSPJRmlCPz\nbljWFdLLTDLtdFtLlwZvqfbgwwHfwqHNvw\nBRRjhRQndRNVqBjRVhFLccjpwMmLmjHmgFHH\nfZJfJvzPPWtWWlltZzZPpcgFMsFFwwFdpHdgwtdw\nPCrdrzzfWCPdvSlqTqNSDnnQVVQQGT\nDjbfBMDSfBljBsLSjSZbzrGtPtMCPtVPvvqrzqzG\nmWdJWcppcNTdpppjzjRRVrPRpq\nQncmnHwmdTmwQcmjNTfgfhlBShshhsffnfbB\nWGDsMJsrjHCWtDMGDDVQqSvZqfSJzSnvnvvv\nLgLFLFBFLVVzfBzMqZ\nlgmFcwLhNcwdwwMLwhmcRDjNpCWRsWRspGGssHCp\nPnPzNccnjFfvCvhbSBVcWqdhSVhV\npsGMDQJDDDJgQNDHHJbwqwBsVqqZVWBBhBdd\nDlDJDQGptpgpGDfTRnrTrFPnNTlf\nMSSSMLLmFHcDScSq\nppZnCsbjPZpnnJcbRDmzHJqRRD\npmNmnGnQNnClZGMVMdBGrMgVWg\nlsTTGcQzBcljCcQzGcGjGptttpmvSJtmggtwwswwtS\nqZRnrhMbRVdhZRhhdnnVRPbmwSNwNNHtmJBvwpvtwNSvSb\nVnMrqrrdqhZrnrBLLlzzlQjQjLfTcGfFDF\ndJJTlHvhZqZlQTJnSgQDzgsSbScsSBzc\nRRNtGjCCpRPPpRtjfrttRzmbscLsLZLgcsbmLzSGLB\nwfNttfNrtWwPNNFfRtpfrdJMTTTZTMZTTVTlVwTlvM\nPQTGLmdNTgPmGgNNdCPLQlrMqBrDzMCMFqDqFqjVCBCD\nhhRwwvpSFmzDrmFh\nvwwZfSfsmvtSspnZLLLdLGWPTGTQtTWG\npMcWzWFvWhFpPMWzvvhpdprHTZTQrHrQdZTJdfTgQTnJ\nCGbjBbNjjDmRHJDgrTVVZg\nNNttGlGqNLsbtlhMFMFcMLwMvvZz\nCGSCBNCQBtBCQttBwCGtGtQrqrLrJqZHLHbqHvLDHLrq\nnVVhPMfVdfVPbfqLLqgDDqPvgZsv\ncpVncbfnhFcBltTplpmTBC\nMrdcdStbMnddtRBdqMnFmbqGCwqCVHVsNHwPfGVPqsCsCs\nDBLllzWWQQzlZVVVCsGWHfsH\nJQphjTgBjlLgjjpTpLgvTjQnnnSJJRRFmdbRRSdMRtmdMc\nQbRZMSWMblwLsgpwZzqZ\nBFncBrfcdNrrnVrNjsFzFTJpJLGJsGqLTp\nVjhDDBdrfdhQMllzHmPQMh\nLdVVjFVFbpVGRQGllG\ncNMcJNHzJWJtCWHNJHcHczWpGmmhMQmBBqrlRhBmpGpGBQ\nJZzTTtCZtHCJnNnNwPfbFpnfdDdLdnvP\nTpMlrWTTddjmlmDmgQgRtw\nMNNVMSsVSNSnNVMFLDqwtGgRRtGbgFRwtR\nCCLSCPSCZZHVCfZscBJJhPphpdpprdhjJM\ngSMSHJHsMMpzRgHzsRMPPSzsPhtZtZdqdDqQDhdCdZmQldht\nFCcCnrGcNTfvvtqqfvlflQ\nTrTrWNWwrTJLMzJCzWLL\nTpTzwMrfbrpFpMbFrrrzbPSdZmtSZRTlTZRlmdCVlCtJ\nvqvWgqDJQJsQCVtZgdZdRRGd\nvsvLJLchWBcqnvczwjLfzPjfrjzPrz\nzqzbqCFZgmzzmNmf\nvpRWSbRVbVWddVpwvwdRSwnSNgLHsnfNgMmgMLMmnrns\nDwWVpJRlpdbpRDWdGJGcGlhFtPPCqCCBFqZPQttlqFBq\nwQRlwtBJBDwttJdGvLfBvHLLfTLz\nMMmNZcMrcMFnRHzfjjvvHfvc\nFggpbFnhrNNrrMrMbMbnhQVJVhstJwqWCVCRsQJQ\nDQbCGblQlpQFQlHjCbjwDQQMggNmJmgnnpRBngfZmNgJMf\nzvhWccWVdWBchdssPrrWZZZfmsmmmgsnZZJRsRTf\nzBdtqPccWPHFCqCCqljq\nttrbRMmgtHgfmHSfBpLfnBBZBppB\nCVTJDCCNPwCPDwcqzmddQZdTQdnLBQThWp\nzwFDjwDJJPzjzVNcVJwCcbRHGmbbMrFHgHvrsgbblG\ngZjjwHqHCzrMZVVR\nhhzcdTzPrVhVCGMb\nfPcmLPNffsccJDdNDjBnpwzmHqgWjHwwvg\nSJQFSvQBlzbSCgdPPddPPPSN\npcrjcWLwwcHcgPNgTPLMNTCB\npRsjsWRnrpHRmrBrHrjlbJFvvzQFnzQblQDDbJ\nVjQVMQPVMfVPPbGPHHbGJD\npcqSttltsbDGddsCJG\nTSchqLtTLFhgQbMMQMrr\ntrqzMRwNTtDzLPJQgWmjmjrf\nlbBQdpZbsmhGmZhmmG\nllVbpCplvvHBBHpnRDcDRRqnRRQnFRzT\nSLSSFFmzLShsVSSHnLnrJdbnRdZZbrRw\nqCfWBftpNWNNlqvTpwrRbGGCnwGmgRJGZn\nNcTBNpvWvBWpMftNffpqWlTpmzPDQPSzFVMsFQVhHsjHszss\nVtJtNBRBGDpdpNbC\nQgLncnttvFcwwhLvFjSGsSbmmQCSDdpCmpdG\nvLgjLhhrctMvLFFjLtMTLMgfPZqBZPZzJBBfWZZPRZZTRV\nmJzDJJpJBvfsGMQnBM\nCwPWCLRRWwRqwPqhPsrZrnrlhhQrMTrvZl\ndCdLLSPRLSqWqVSLqLjgJDzDmtbngFVtJtzz\nmtgWtMWrqjzQTTjghwwfczlNJdlcJnlc\nFvRsDPPFGRBFvvslwDnTlcTTdwndlh\nSGBZRBTsFGBRvLpvSCmgQWQjgggMrQjmmSmW\nGcsRrQhrVVjhRcWlnDFGGmvntDWZ\nTPbSgJJgBSCbCTbLHMCMTTZdFHvtZlWZDZFzmzZHZmmF\ngBCMCSpbPMMPjcjqQQpqQprv\nnZJcnZwvwzvTTTVtpDFnHH\nDQPBqGGGdMdTRHRBpNgFNR\ndCGPfhPWQdWWWCWShWPqrChWLLwLswjcvSJbvbLjJLbzJbJD\nQrBQtdtrQBrdtFHPrdQBDvGhLGnPnCWnmpDmLpmD\nNjlRJRlNzJJVbSSRVZwwJcmpWDGCWnbchnLCCmnWCG\nllSJzsZzMMlsSZjSjZwJNQqtHHdBFsqdfTHhqFftQB\nzdTJFHTdDBzrNdMnhNnNdM\nZlLZZcLtVtcWtGjtzLjLZjCrnVNrnRbrQQbQSRVrRnSNqS\nlZtGtCvjZPCGCctPpsDDBzTHFmPmFszD\nmQSMvdMQtQdZhQrPWCPqPQrN\nRwjwnZGzJFTZgzggzJDDwJnCPPhNNqPrLhrGNcWcWNPqCq\nZTzDfnwFzTngTwJvfSlMtMMlmsHmHt\nlZlmFRVZWmgQWhRsRpJsCJpJct\nPTbPTGTGwwGrbdfjNNZJvcCsCZtvpTsh\nbGdBBqGrdBPjDMzzVFZgqQzFFL\nszvsmLvppPPtzGLGWpVdTSHTNgjHQRmHTgSH\nFnBMBNZwZNcnDZMcnZlZgwgdQTTHjVJjHHVRQHJj\nDnZrFCMZMNffrLPbLsfW\nrJvmnBgnrCrGRSGNQR\nhthjNfhwctwpjTLtVLjTGSpldSCGSPdlPSRzSqSz\nTVcTfHNFcwtjMhTvgbHZsBbWmmZbnH\nWsQgstQmvQJnssWsWPzhRzhBjZBSBRZSnj\nqwCNqFwDrrlDrFPvRhTSPPzLRz\nbppqwppCddlvfbDNVgmMmtMfVVmfmVWW"
	utils.Run(solve, testInput, true)
	utils.Run(solve, realInput, true)
}
