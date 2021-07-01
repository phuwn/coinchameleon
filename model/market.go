package model

import "time"

type MarketData struct {
	PairID string     `json:"pair_id,omitempty" sql:"pair_id"`
	TS     *time.Time `json:"ts,omitempty" sql:"ts"`
	Open   string     `json:"open,omitempty" sql:"open"`
	High   string     `json:"high,omitempty" sql:"high"`
	Low    string     `json:"low,omitempty" sql:"low"`
	Close  string     `json:"close,omitempty" sql:"close"`
	Volume string     `json:"volume,omitempty" sql:"volume"`
}

func (m MarketData) TableName() string {
	return "market"
}
