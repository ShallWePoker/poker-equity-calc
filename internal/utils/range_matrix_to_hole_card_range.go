package utils

import (
	"errors"
	"fmt"
	"github.com/ShallWePoker/poker-equity-calc/internal/consts"
	"github.com/ShallWePoker/poker-equity-calc/internal/models"
)

func RangeMatrixToHoleCardRanges(matrix [][]int) ([]models.HoleCardRange, error) {
	holeCardRanges := make([]models.HoleCardRange, 0)
	if len(matrix) > 13*13 {
		return holeCardRanges, errors.New("matrix must have not more than 13*13 entries")
	}
	for i, entry := range matrix {
		if len(entry) != 2 {
			return holeCardRanges, errors.New("each matrix entry must have 2 coordinates exactly")
		}
		x := entry[0]
		y := entry[1]
		holeCardRange := models.HoleCardRange{}
		titlePair := ""
		if title, exists := consts.MatrixCoordinateToTitle[x]; !exists {
			return holeCardRanges, errors.New(fmt.Sprintf("matrix #%dth entry %v x coordinate %d must be in [0, 12]", i, entry, x))
		} else {
			titlePair = titlePair+title
		}
		if title, exists := consts.MatrixCoordinateToTitle[y]; !exists {
			return holeCardRanges, errors.New(fmt.Sprintf("matrix #%dth entry %v y coordinate %d must be in [0, 12]", i, entry, y))
		} else {
			titlePair = titlePair+title
		}
		holeCardRange.TitlePair = titlePair
		holeCardRange.Suited = x > y
		holeCardRanges = append(holeCardRanges, holeCardRange)
	}
	return holeCardRanges, nil
}