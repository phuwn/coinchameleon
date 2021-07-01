package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/phuwn/coinchameleon/data/db"
	"github.com/phuwn/coinchameleon/util"
)

var (
	pairList  = [...]string{"BTCUSDT", "ETHUSDT", "BNBUSDT", "DOGEUSDT", "FTMUSDT", "MATICUSDT"}
	insertSQL = `INSERT INTO MARKET (PAIR_ID,TS,OPEN_PRICE,HIGH_PRICE,LOW_PRICE,CLOSE_PRICE,VOLUME) VALUES ('%v',TO_DATE('%v','DD/MM/YYYY'),%v,%v,%v,%v,%v)`
)

func crawl() (err error) {

	tx := db.Get().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	startTime, err := time.Parse("02/01/2006 15:04:05", "01/01/2020 00:00:00")
	if err != nil {
		return err
	}

	apiKey, secretKey := util.Getenv("BINANCE_APIKEY", ""), util.Getenv("BINANCE_SECRETKEY", "")
	client := binance.NewClient(apiKey, secretKey)

	for _, pair := range pairList {
		klineData, err := client.NewKlinesService().
			Symbol(pair).
			StartTime(startTime.UTC().Unix() * 1000).
			Interval("1d").
			Limit(1000).
			Do(context.Background())
		if err != nil {
			return err
		}

		log.Printf("[CRAWL] %v data of %v succeed\n", len(klineData), pair)

		for _, v := range klineData {

			q := fmt.Sprintf(
				insertSQL,
				pair,
				time.Unix(v.OpenTime/1000, 0).Format("02/01/2006"),
				v.Open,
				v.Close,
				v.High,
				v.Low,
				v.Volume,
			)

			if err := tx.Exec(q).Error; err != nil {
				return err
			}
		}

		log.Printf("[STORE] data of %v succeed\n---\n", pair)
	}
	return nil
}
