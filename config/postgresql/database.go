package postgresql

import (
	"fmt"
	"os"
	"strconv"
	"time"

	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/zerolog/log"
	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConnection struct {
	// DbPayment *gorm.DB
	Db *gorm.DB
}

func CreateConnection() *DbConnection {
	log.Debug().Msg(os.Getenv("DRIVER_NAME_DB"))

	var db *gorm.DB
	var err error
	i := 0
	for {

		db, err = gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING_DB")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			if i == 0 {
				fmt.Printf("CONNECT %v : %v \n", err, os.Getenv("CONNECTION_STRING_DB"))
			} else {
				fmt.Printf("RECONNECT(%d) %v : %v \n", i, err, os.Getenv("CONNECTION_STRING_DB"))
			}
			time.Sleep(3 * time.Second)
			i++
			continue
		}
		break
	}

	num, _ := strconv.Atoi(os.Getenv("MAX_CONNECTION_POLL_DB"))
	dbSQL, err := db.DB()
	if err != nil {
		log.Error().Msg(err.Error())
		log.Info().Msg("failed to connect database")
	}
	maxConnIdle, err := strconv.Atoi(os.Getenv("MAX_CONNECTION_IDLE"))
	if err != nil {
		maxConnIdle = 5
	}
	dbSQL.SetConnMaxIdleTime(time.Duration(maxConnIdle))
	dbSQL.SetMaxOpenConns(num)
	dbSQL.SetConnMaxLifetime(time.Hour)
	return &DbConnection{db}
}
