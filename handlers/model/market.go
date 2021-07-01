package model

import "time"

type MarketData struct {
	CoinID      string     `json:"coin_id,omitempty" sql:"coin_id"`
	TS          *time.Time `json:"ts,omitempty" sql:"ts"`
	Price       float32    `json:"price,omitempty" sql:"price"`
	MarketCap   float32    `json:"market_cap,omitempty" sql:"market_cap"`
	TotalVolume float32    `json:"total_volume,omitempty" sql:"total_volume"`
}

func (m MarketData) TableName() string {
	return "market"
}

type PriceData struct {
	Price float32 `json:"price,omitempty" sql:"price"`
	Date  string  `json:"date,omitempty" sql:"date"`
}
