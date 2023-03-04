package db

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"test-bpjs/src/api/v1/models"
	"time"
)

type Database struct {
	DB *gorm.DB
}

func NewDB(conf models.ServerConfig) *Database {
	var DB *gorm.DB
	var err error

	var host, user, password, name, port string

	host = conf.DBConfig.Host
	port = conf.DBConfig.Port
	user = conf.DBConfig.User
	password = conf.DBConfig.Password
	name = conf.DBConfig.Name

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, name, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Jakarta")
			return time.Now().In(ti)
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	var ctx context.Context

	DB = DB.WithContext(ctx)

	dbSQL, err := DB.DB()
	if err != nil {
		log.Println("config/DB: gorm open connect")
		log.Fatal(err)
	}

	//Database Connection Pool
	dbSQL.SetMaxIdleConns(10)
	dbSQL.SetMaxOpenConns(100)
	dbSQL.SetConnMaxLifetime(time.Hour)

	err = dbSQL.Ping()
	if err != nil {
		log.Println("config/DB: can't ping the DB, WTF")
		log.Fatal(err)
	}
	go doEvery(10*time.Minute, pingDb, DB)

	models.InitTable(DB)

	return &Database{
		DB: DB,
	}

}

func doEvery(d time.Duration, f func(*gorm.DB), y *gorm.DB) {
	for _ = range time.Tick(d) {
		f(y)
	}
}

func pingDb(db *gorm.DB) {
	log.Println("PING CONNECTION")
	dbSQL, err := db.DB()
	if err != nil {
		log.Println("PING CONNECTION FAILURE")
	}

	err = dbSQL.Ping()
	if err != nil {
		log.Println("PING CONNECTION FAILURE")
	}
}
