package models

import "fmt"

type HoleCardEquity struct {
	HoleCard HoleCard
	Equity   float64
	WinRate   float64
	TieRate   float64
}

func (hce HoleCardEquity) ToString() string {
	return fmt.Sprintf("holecard %s equity: Equity %.3f%%, WinRate %.3f%%, TieRate %.3f%%", hce.HoleCard.ToString(), hce.Equity*100, hce.WinRate*100, hce.TieRate*100)
}
