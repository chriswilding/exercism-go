package queenattack

import (
	"errors"
	"math"
)

var InvalidRankOrFile = errors.New("invalid rank or file")

func CanQueenAttack(white, black string) (bool, error) {
	if white == black || !valid(white) || !valid(black) {
		return false, InvalidRankOrFile
	}

	if white[0] == black[0] || white[1] == black[1] {
		return true, nil
	}

	file := float64(int(white[0]) - int(black[0]))
	rank := float64(int(white[1]) - int(black[1]))

	cqa := math.Abs(file) == math.Abs(rank)
	return cqa, nil
}

func valid(queen string) bool {
	return len(queen) == 2 && queen[0] >= 'a' && queen[0] <= 'h' && queen[1] >= '1' && queen[1] <= '8'
}
