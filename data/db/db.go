package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cengsin/oracle"
	"gorm.io/gorm"
)

var (
	db *gorm.DB

	ErrMissingTx = fmt.Errorf("missing transaction in context")
)

func newDB(connectionInfo string) error {
	newDB, err := gorm.Open(oracle.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		return err
	}
	newDB.NowFunc = func() time.Time { return time.Now().UTC() }

	db = newDB
	return nil
}

func Get() *gorm.DB {
	return db
}

func Start() {
	err := newDB(os.Getenv("DATASOURCE"))
	if err != nil {
		log.Fatal("Error connect db :", err)
	}
}

func Close() error {
	con, err := db.DB()
	if err != nil {
		return err
	}
	return con.Close()
}

func Healthz() error {
	con, err := db.DB()
	if err != nil {
		return err
	}
	return con.Ping()
}
