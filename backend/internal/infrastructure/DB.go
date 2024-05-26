package infrastructure

import (
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbonce sync.Once

var rs *redis.Client
var redisonce sync.Once

func GetDB() *gorm.DB {
	dbonce.Do(func() {
		dbstr := os.Getenv("DB_CONNECTION")
		var err error
		db, err = gorm.Open(
			mysql.Open(dbstr),
			&gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
	return db
}

func GetRedis() *redis.Client {
	redisonce.Do(func() {
		redisPass := os.Getenv("REDIS_PASS")
		rs = redis.NewClient(&redis.Options{
			Addr:     "45.147.97.39:6379",
			Password: redisPass,
		})
	})
	return rs
}
