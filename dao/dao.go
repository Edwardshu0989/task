package dao

import (
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var Provider = wire.NewSet(New, NewRedis)

type Dao struct {
	db *gorm.DB
	rd *redis.Client
}

func NewDB(db *gorm.DB, redis *redis.Client) (d *Dao) {
	d = &Dao{
		db: db,
		rd: redis,
	}
	return
}
