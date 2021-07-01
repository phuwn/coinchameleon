package handlers

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/phuwn/coinchameleon/data/db"
	"github.com/phuwn/coinchameleon/handlers/model"
)

func getPriceData(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(400, "missing id")
	}
	date := c.QueryParam("date")
	if date == "" {
		return c.String(400, "date is missing")
	}
	tx := db.Get()

	var priceData []*model.PriceData

	err := tx.Raw(
		fmt.Sprintf(
			`SELECT price, TO_CHAR(ts, 'DD MON') as "DATE" from market PARTITION(prices_%v) where coin_id = '%s'`,
			date,
			id,
		),
	).Find(&priceData).Error
	if err != nil {
		return c.String(404, err.Error())
	}

	return c.JSON(200, &priceData)
}
