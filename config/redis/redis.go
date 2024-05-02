package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	zerolog "github.com/rs/zerolog/log"
)

type RedisConnection interface {
	// redis communication
	Get(ctx echo.Context, key, field string) ([]byte, error)
	Set(ctx echo.Context, key, field string, value interface{}, expDur time.Duration) error
	Del(ctx echo.Context, key, field string)

	// redis implementation
	CreateCache(ctx echo.Context, key, field string, value interface{}, dur time.Duration)
	GetCache(ctx echo.Context, key string, field string, res interface{}) bool
	DeleteCache(ctx echo.Context, key string, field string)
}

type redisCtx struct {
	redisClient redis.Cmdable
}

func CreateConnection() RedisConnection {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	poolSize, _ := strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
	minConnIdle, _ := strconv.Atoi(os.Getenv("REDIS_CONN_MIN_IDLE"))
	convertedTimeout, _ := strconv.Atoi(os.Getenv("REDIS_TIMEOUT"))
	timeout := time.Duration(convertedTimeout) * time.Second
	client := redis.NewClient(&redis.Options{
		DB:           db,
		Addr:         fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password:     os.Getenv("REDIS_PASSWORD"),
		PoolSize:     poolSize,
		PoolTimeout:  timeout,
		MinIdleConns: minConnIdle,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		zerolog.Error().Msg(err.Error())
		// zerolog.Panic().Msg("failed connect to redis")
	}

	zerolog.Debug().Msg(fmt.Sprintf("Redis: %s", pong))
	fmt.Println("check redis")
	return &redisCtx{
		redisClient: client,
	}
}

//-- start of redis communication
// redis HSet
func (c *redisCtx) Set(ctx echo.Context, key, field string, value interface{}, dur time.Duration) error {
	payload, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = c.redisClient.HSet(key, field, payload).Err()
	if err != nil {
		zerolog.Error().Err(err).Msg("error redis HSet")
	}

	if dur > 0 {
		err = c.redisClient.Expire(key, dur).Err()
		if err != nil {
			zerolog.Error().Err(err).Msg("error redis Set Expire")
		}
	}

	zerolog.Debug().Msg(fmt.Sprintf("Redis: set %s key on %s field operation success.", key, field))

	return nil
}

// redis HGet
func (c *redisCtx) Get(ctx echo.Context, key, field string) ([]byte, error) {
	data, err := c.redisClient.HGet(key, field).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			zerolog.Error().Err(err).Msg("error redis HGet fetch")
		}
		return nil, err
	}
	zerolog.Debug().Msg(fmt.Sprintf("Redis: get %s key on %s field operation success.", key, field))
	return []byte(data), nil
}

// redis HDel
func (c *redisCtx) Del(ctx echo.Context, key, field string) {
	err := c.redisClient.HDel(key, field).Err()
	if err != nil {
		zerolog.Error().Err(err).Msg("error redis HDel")
	}
	zerolog.Debug().Msg(fmt.Sprintf("Redis: del %s key on %s field operation success.", key, field))
}

//-- end of redis communication

//-- start of redis implementation
func (r *redisCtx) GetCache(ctx echo.Context, key string, field string, res interface{}) bool {
	var fromRedis bool
	dataByte, err := r.Get(ctx, key, field)
	if err != nil && err != redis.Nil {
		zerolog.Error().Err(err).Str("key", key).Str("field", field).Msg("error redis fetch HGet")
	}

	if err == nil && dataByte != nil {
		if err := json.Unmarshal(dataByte, &res); err != nil {
			zerolog.Error().Err(err).Msg("error unmarshal")
		} else {
			fromRedis = true
		}
	}

	return fromRedis
}

func (r *redisCtx) CreateCache(ctx echo.Context, key, field string, value interface{}, dur time.Duration) {
	if err := r.Set(ctx, key, field, value, dur); err != nil {
		zerolog.Error().Err(err).Msg("error HSet")
	}
}

func (r *redisCtx) DeleteCache(ctx echo.Context, key, field string) {
	r.Del(ctx, key, field)
}

//-- end of redis implementation
