package dao

import (
	"awesomeProject/config"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/prometheus/common/log"
	"time"
)

func NewRedis() (r *redis.Client, cf func(), err error) {
	r = redis.NewClient(&redis.Options{
		Addr:     config.Env.RedisAddr,
		Password: "",
		DB:       0,
	})
	cf = func() { r.Close() }
	_, err = r.Ping().Result()
	return
}

func (d *Db) AddData() bool {
	if _, err := d.rd.Set("test-redis", "111", 10*time.Minute).Result(); err != nil {
		log.Info(fmt.Sprintf("%s", "redis设置错误"))
		return false
	}
	return true
}

func (d *Db) GetData() (string, error) {
	val, err := d.rd.Get("test-redis").Result()
	if err != nil {
		log.Info("查询验证码出错")
		return "", err
	}
	return val, nil
}
