package queenattack

import (
	"errors"
	"math"
)

var InvalidRankOrFile = errors.New("invalid rank or file")

func CanQueenAttack(white, black string) (bool, error) {
	if white == black || len(white) != 2 || len(black) != 2 {
		return false, InvalidRankOrFile
	}

	wr, wf := queenToRankAndFile(white)
	br, bf := queenToRankAndFile(black)

	if !validRankOrFile(wf) || !validRankOrFile(wr) || !validRankOrFile(bf) || !validRankOrFile(br) {
		return false, InvalidRankOrFile
	}

	if wf == bf || wr == br {
		return true, nil
	}

	cqa := math.Abs(float64(wf-bf)) == math.Abs(float64(wr-br))
	return cqa, nil
}

func queenToRankAndFile(queen string) (int, int) {
	file := int(queen[0] - 96) // convert a to h to an integer
	rank := int(queen[1] - 48) // convert 1 to 9 to an integer
	return rank, file
}

func validRankOrFile(rOrF int) bool {
	return 0 < rOrF && rOrF < 9
}
