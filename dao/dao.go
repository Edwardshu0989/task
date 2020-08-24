package dao

import (
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Provider = wire.NewSet(New, NewRedis)

type Db struct {
	db *gorm.DB
	rd *redis.Client
}

func NewDB(db *gorm.DB, redis *redis.Client) (d *Db) {
	d = &Db{
		db: db,
		rd: redis,
	}
	return
}
