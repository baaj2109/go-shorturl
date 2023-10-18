package common

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Redis *redis.Client

func InitDB() {
	InitMysql()
	InitRedis()
}

func InitMysql() (*sql.DB, error) {
	username := Config.GetString("mysql.username")
	password := Config.GetString("mysql.password")
	host := Config.GetString("mysql.host")
	port := Config.GetInt("mysql.port")
	dbname := Config.GetString("mysql.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(Config.GetInt("mysql.max-open-conns"))
	db.SetMaxIdleConns(Config.GetInt("mysql.max-idle-conns"))
	duration, err := time.ParseDuration(Config.GetString("mysql.max-open-time"))
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitRedis() {
	addr := Config.GetString("redis.addr") + ":" + Config.GetString("redis.port")
	password := Config.GetString("redis.passwd")

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
}
