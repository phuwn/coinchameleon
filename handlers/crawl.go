package handlers

import (
	"log"
	"time"

	"github.com/phuwn/coinchameleon/data/db"
	"github.com/phuwn/coinchameleon/handlers/model"
	coingecko "github.com/superoo7/go-gecko/v3"
)

var (
	usd       = "usd"
	timerange = "max"
	coinList  = [...]string{"bitcoin", "dogecoin", "ethereum", "binancecoin", "matic-network"}
)

func CrawlMarket(cg *coingecko.Client) {
	tx := db.Get()

	for _, c := range coinList {
		d, err := cg.CoinsID(c, false, false, false, false, false, false)
		if err != nil {
			log.Fatal(err)
		}
		if err := tx.Create(&model.Coin{c, d.Name, d.Symbol}).Error; err != nil {
			log.Fatal(err)
		}
	}

	for _, coin := range coinList {
		chart, err := cg.CoinsIDMarketChart(coin, usd, timerange)
		if err != nil {
			log.Fatal(err)
		}

		for i, v := range *chart.Prices {
			ts := time.Unix(int64(v[0]/1000), 0)
			if err := tx.Create(&model.MarketData{
				coin,
				&ts,
				v[1],
				(*chart.MarketCaps)[i][1],
				(*chart.TotalVolumes)[i][1],
			}).Error; err != nil {
				log.Fatal(err)
			}
		}
	}
}
