package main

import "fmt"

var _seq = []int{5, 1, 22, 25, 6, -1, 8, 10}
var _matchSeq = []int{1, 6, -1, 10}

func VarSeqMatch(_seq []int, _matchSeq []int) bool {
	arrIdx := 0
	seqIdx := 0

	for arrIdx < len(_seq) && seqIdx < len(_matchSeq) {
		if _seq[arrIdx] == _matchSeq[seqIdx] {
			seqIdx += 1
		}
		arrIdx += 1
	}
	return seqIdx == len(_matchSeq)
}

func main() {
	res := VarSeqMatch(_seq, _matchSeq)
	fmt.Println(res)
}
