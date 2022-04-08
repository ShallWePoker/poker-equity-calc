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

func InitHoleCardRange(titlePair string, suited bool) (HoleCardRange, error) {
	if len(titlePair) != 2 {
		return HoleCardRange{}, errors.New("titlePair must have 2 str")
	}
	titlePair = strings.ToUpper(titlePair)
	_, ok := ValidTitles[string(titlePair[0])]
	if !ok {
		return HoleCardRange{}, errors.New(fmt.Sprintf("Invalid titlePair: %s ", string(titlePair[0])))
	}
	_, ok = ValidTitles[string(titlePair[1])]
	if !ok {
		return HoleCardRange{}, errors.New(fmt.Sprintf("Invalid titlePair: %s ", string(titlePair[1])))
	}
	hcr := HoleCardRange{TitlePair: titlePair, Suited: suited}
	return hcr, nil
}
