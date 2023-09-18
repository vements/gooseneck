package gooseneck

import (
	"strconv"

	"github.com/redis/go-redis/v9"
)

const (
	REDIS_ADDR     = "REDIS_ADDR"
	REDIS_DATABASE = "REDIS_DATABASE"
	REDIS_USERNAME = "REDIS_USERNAME"
	REDIS_PASSWORD = "REDIS_PASSWORD"
)

type RedisEnv struct {
	Env
}

func (e RedisEnv) Addr() string {
	return e.MustDefine(REDIS_ADDR)
}

func (e RedisEnv) Database() int {
	if db, err := strconv.Atoi(e.Optional(REDIS_DATABASE, "0")); err != nil {
		panic(err)
	} else {
		return db
	}
}

func (e RedisEnv) Username() string {
	return e.MustDefine(REDIS_USERNAME)
}

func (e RedisEnv) Password() string {
	return e.MustDefine(REDIS_PASSWORD)
}

func NewRedisClient() *redis.Client {
	env := RedisEnv{}
	Info().Str("host", env.Addr()).Msg("using redis")
	return redis.NewClient(&redis.Options{
		Addr:     env.Addr(),
		Password: env.Username(),
		Username: env.Password(),
		DB:       env.Database(),
	})
}
