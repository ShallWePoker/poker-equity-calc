package consts

const (
	PREFLOP = "preflop"
	FLOP    = "flop"
	TURN    = "turn"
	RIVER   = "river"
)

var STAGE_UNKNOWN_CARD_NUM = map[string]int {
	PREFLOP: 5,
	FLOP: 2,
	TURN: 1,
	RIVER: 0,
}