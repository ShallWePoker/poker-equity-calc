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
		holeCardRange = HoleCardRange{TitlePair: string(titlePair[1]) + string(titlePair[0]), Suited: suited}
	}
	holeCardRange = HoleCardRange{TitlePair: titlePair, Suited: suited}
	return holeCardRange, nil
}

func (r HoleCardRange) ToString() string {
	if r.TitlePair[0] == r.TitlePair[1] {
		return r.TitlePair
	}
	if r.Suited {
		return r.TitlePair + "s"
	} else {
		return r.TitlePair + "o"
	}
}

type UnbalancedHolecardRange struct {
	HolecardRanges []HoleCardRange
	Percentage     float64
}

type UnbalancedHolecardRanges []UnbalancedHolecardRange

func (u UnbalancedHolecardRanges) ToString() string {
	str := ""
	for _, ubRange := range u {
		for _, ubRangePart := range ubRange.HolecardRanges {
			str = str + fmt.Sprintf("%s ", ubRangePart.ToString())
		}
		str = str + fmt.Sprintf("%0.2f%%;", ubRange.Percentage*float64(100))
	}
	return str
}