package conn

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mojo-zd/helm-crabstick/data/types"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host        string
	Port        string
	Username    string
	Password    string
	Schema      string
	MaxConn     int
	MaxIdleConn int
	MaxIdleTime int
}

var db *gorm.DB

func NewConn(cfg Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Schema,
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)
	mydb, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db, err = gorm.Open(mysql.New(mysql.Config{Conn: mydb}), &gorm.Config{Logger: newLogger})
	sql, err := db.DB()
	if err != nil {
		panic(err)
	}

	sql.SetMaxOpenConns(cfg.MaxConn)
	sql.SetConnMaxIdleTime(time.Duration(cfg.MaxIdleTime))
	sql.SetMaxIdleConns(cfg.MaxIdleConn)
	if err = autoMigrate(); err != nil {
		panic(err)
	}
}

func autoMigrate() error {
	return db.AutoMigrate(&types.Release{})
}

func GetDB() *gorm.DB {
	return db
}
