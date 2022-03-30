package models

type Card struct {
	Title string   `json:"title"`
	Suit  string   `json:"suit"`
	Rank  int      `json:"rank"`
}

func (card Card) ToString() string {
	return card.Title+card.Suit
}