package models

import (
	"fmt"
	"strings"
)

type HoleCardEquity struct {
	HoleCard HoleCard
	Equity   float64
	WinRate   float64
	TieRate   float64
}

func (hce HoleCardEquity) ToString() string {
	return fmt.Sprintf("holecard %s equity: Equity %.3f%%, WinRate %.3f%%, TieRate %.3f%%", hce.HoleCard.ToString(), hce.Equity*100, hce.WinRate*100, hce.TieRate*100)
}

type HoleCardRangeEquity struct {
	HoleCardRange []HoleCardRange
	Equity   float64
	WinRate   float64
	TieRate   float64
}

func (rgeEquity HoleCardRangeEquity) ToString() string {
	holecardrange := ""
	for _, rge := range rgeEquity.HoleCardRange {
		holecardrange = holecardrange+rge.ToString()+","
	}
	holecardrange = strings.TrimSuffix(holecardrange, ",")
	return fmt.Sprintf("range %s equity: Equity %.3f%%, WinRate %.3f%%, TieRate %.3f%%", holecardrange, rgeEquity.Equity*100, rgeEquity.WinRate*100, rgeEquity.TieRate*100)
}