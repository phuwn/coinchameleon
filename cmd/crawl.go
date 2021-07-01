package main

import (
	"context"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/phuwn/coinchameleon/data/db"
	"github.com/phuwn/coinchameleon/model"
	"github.com/phuwn/coinchameleon/util"
)

var pairList = [...]string{"BTCUSDT", "ETHUSDT", "BNBUSDT", "DOGEUSDT", "FTMUSDT", "MATICUSDT"}

func crawl() (err error) {

	tx := db.Get().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	startTime, err := time.Parse("02/01/2006 15:04:05", "01/01/2021 00:00:00")
	if err != nil {
		return err
	}

	apiKey, secretKey := util.Getenv("BINANCE_APIKEY", ""), util.Getenv("BINANCE_SECRETKEY", "")
	client := binance.NewClient(apiKey, secretKey)

	for _, pair := range pairList {
		klineData, err := client.NewKlinesService().
			Symbol("BTCUSDT").
			StartTime(startTime.UTC().Unix() * 1000).
			Interval("1d").
			Limit(1000).
			Do(context.Background())
		if err != nil {
			return err
		}

		for _, v := range klineData {
			ts := time.Unix(v.OpenTime/1000, 0)
			newRecord := &model.MarketData{
				pair,
				&ts,
				v.Open,
				v.High,
				v.Low,
				v.Close,
				v.Volume,
			}

			if err := tx.Create(&newRecord).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
