package common

import (
	conf "blog-go-gin/config"
	"blog-go-gin/logging"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	redisClient *redis.Client
)

var ctx = context.Background()

type redisUtil struct{}

func InitRedis() {
	config := conf.GetConf()
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.Redis.RedisConn,
		Password: config.Redis.RedisPwd,
		DB:       config.Redis.Db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		logging.Logger.Errorf("connect to redis failed, err: %s", err)
		return
	}
}

func (*redisUtil) Get(key string) (string, error) {
	val, err := redisClient.Get(ctx, key).Result()
	logging.Logger.Errorf("redisUtil get err:%v", err)
	if err != nil {
		return "", err
	}
	return val, nil
}
func (*redisUtil) SetEx(key string, value string, expiration int64, unit time.Duration) {
	err := redisClient.SetEX(ctx, key, value, time.Duration(expiration)*unit).Err()
	if err != nil {
		panic(err)
	}
}

func (*redisUtil) SAdd(key string, member interface{}) {
	redisClient.SAdd(ctx, key, member)
}

func (*redisUtil) SMembers(key string) []string {
	setList, err := redisClient.SMembers(ctx, key).Result()
	if err != nil {
		return []string{}
	}
	return setList
}

func (*redisUtil) SRems(key string, members ...interface{}) {
	redisClient.SRem(ctx, key, members)
}

func (*redisUtil) Delete(key string) {
	redisClient.Del(ctx, key)
}

func (*redisUtil) Keys(pattern string) []string {
	val, err := redisClient.Keys(ctx, pattern).Result()
	if err != nil {
		return []string{}
	}
	return val
}

func (*redisUtil) MultiGet(keys []string) []interface{} {
	val, err := redisClient.MGet(ctx, keys...).Result()
	if err != nil {
		return []interface{}{}
	}
	return val
}

func (*redisUtil) MultiDelete(key []string) {
	redisClient.Del(ctx, key...)
}

func (*redisUtil) Set(key string, value string) error {
	err := redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (*redisUtil) IncrBy(key string, increment int64) int64 {
	val, err := redisClient.IncrBy(ctx, key, increment).Result()
	if err != nil {
		panic(err)
	}
	return val
}

// HashSet 向key的hash中添加元素field的值
func (*redisUtil) HashSet(key, field string, data interface{}) error {
	err := redisClient.HSet(ctx, key, field, data).Err()
	if err != nil {
		return err
	}
	return nil
}

// BatchHashSet 批量向key的hash添加对应元素field的值
func (*redisUtil) BatchHashSet(key string, fields map[string]interface{}) (bool, error) {
	val, err := redisClient.HMSet(ctx, key, fields).Result()
	if err != nil {
		logging.Logger.Error("Redis HMSet Error:", err)
		return false, err
	}
	return val, nil
}

// HashGet 通过key获取hash的元素值
func (*redisUtil) HashGet(key, field string) (string, error) {
	result := ""
	val, err := redisClient.HGet(ctx, key, field).Result()
	if err == redis.Nil {
		logging.Logger.Info("Key Doesn't Exists:", field)
		return result, err
	} else if err != nil {
		logging.Logger.Error("Redis HGet Error:", err)
		return result, err
	}
	return val, nil
}

// BatchHashGet 批量获取key的hash中对应多元素值
func (*redisUtil) BatchHashGet(key string, fields ...string) (map[string]interface{}, error) {
	resMap := make(map[string]interface{})
	for _, field := range fields {
		var result interface{}
		val, err := redisClient.HGet(ctx, key, fmt.Sprintf("%s", field)).Result()
		if err == redis.Nil {
			logging.Logger.Info("Key Doesn't Exists:", field)
			resMap[field] = result
		} else if err != nil {
			logging.Logger.Error("Redis HMGet Error:", err)
			return nil, err
		}
		if val != "" {
			resMap[field] = val
		} else {
			resMap[field] = result
		}
	}
	return resMap, nil
}

func (*redisUtil) HashGetAll(key string) (map[string]string, error) {
	result, err := redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		logging.Logger.Error("Redis HashGetAll Error:", err)
		return nil, err
	}
	return result, nil
}

var RedisUtil = &redisUtil{}
