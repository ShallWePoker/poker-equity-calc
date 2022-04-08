package models

import (
	"errors"
	"fmt"
	"strings"
)

type HoleCardRange struct {
	TitlePair string
	Suited    bool
}

func InitHoleCardRange(titlePair string, suited bool) (holeCardRange HoleCardRange, err error) {
	if len(titlePair) != 2 {
		return HoleCardRange{}, errors.New(fmt.Sprintf("titlePair [%s] must have 2 letters", titlePair))
	}
	titlePair = strings.ToUpper(titlePair)
	rank0, ok := ValidTitles[string(titlePair[0])]
	if !ok {
		return HoleCardRange{}, errors.New(fmt.Sprintf("Invalid title letter: %s ", string(titlePair[0])))
	}
	rank1, ok := ValidTitles[string(titlePair[1])]
	if !ok {
		return HoleCardRange{}, errors.New(fmt.Sprintf("Invalid title letter: %s ", string(titlePair[1])))
	}
	if rank0 <= rank1 {
		holeCardRange = HoleCardRange{TitlePair: string(titlePair[1])+string(titlePair[0]), Suited: suited}
	}
	holeCardRange = HoleCardRange{TitlePair: titlePair, Suited: suited}
	return holeCardRange, nil
}
