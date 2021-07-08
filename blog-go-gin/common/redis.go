package common

import (
	conf "blog-go-gin/config"
	"blog-go-gin/logging"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var ctx = context.Background()
var _instance *RedisUtil

type RedisUtil struct {
	sync.RWMutex
	redisClient *redis.Client
}

func NewRedisUtil() *RedisUtil {
	return &RedisUtil{redisClient: nil}
}

func init() {
	_instance = NewRedisUtil()
	_instance.InitRedis()
}

func GetRedisUtil() *RedisUtil {
	return _instance
}
func (r *RedisUtil) RedisClient() *redis.Client {
	r.RLock()
	rdc := r.redisClient
	r.RUnlock()
	return rdc
}

func (r *RedisUtil) InitRedis() {
	config := conf.GetConf()
	r.Lock()
	r.redisClient = redis.NewClient(&redis.Options{
		Addr:     config.Redis.RedisConn,
		Password: config.Redis.RedisPwd,
		DB:       config.Redis.Db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	r.Unlock()
	defer cancel()

	_, err := r.redisClient.Ping(ctx).Result()
	if err != nil {
		logging.Logger.Errorf("connect to redis failed, err: %s", err)
		return
	}
}

func (r *RedisUtil) Get(key string) (string, error) {
	val, err := r.redisClient.Get(ctx, key).Result()
	logging.Logger.Errorf("redisUtil get err:%v", err)
	if err != nil {
		return "", err
	}
	return val, nil
}
func (r *RedisUtil) SetEx(key string, value string, expiration int64, unit time.Duration) {
	err := r.redisClient.SetEX(ctx, key, value, time.Duration(expiration)*unit).Err()
	if err != nil {
		panic(err)
	}
}

func (r *RedisUtil) SAdd(key string, member interface{}) {
	r.redisClient.SAdd(ctx, key, member)
}

func (r *RedisUtil) SMembers(key string) ([]string, error) {
	setList, err := r.redisClient.SMembers(ctx, key).Result()
	if err != nil {
		return []string{}, err
	}
	return setList, nil
}

func (r *RedisUtil) SIsMember(key, member string) (bool, error) {
	isExist, err := r.redisClient.SIsMember(ctx, key, member).Result()
	if err != nil {
		return false, err
	}
	return isExist, nil
}

func (r *RedisUtil) SCard(key string) (int64, error) {
	result, err := r.redisClient.SCard(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (r *RedisUtil) SRems(key string, members ...interface{}) {
	r.redisClient.SRem(ctx, key, members)
}

func (r *RedisUtil) Delete(key string) {
	r.redisClient.Del(ctx, key)
}

func (r *RedisUtil) Keys(pattern string) []string {
	val, err := r.redisClient.Keys(ctx, pattern).Result()
	if err != nil {
		return []string{}
	}
	return val
}

func (r *RedisUtil) MultiGet(keys []string) []interface{} {
	val, err := r.redisClient.MGet(ctx, keys...).Result()
	if err != nil {
		return []interface{}{}
	}
	return val
}

func (r *RedisUtil) MultiDelete(key []string) {
	r.redisClient.Del(ctx, key...)
}

func (r *RedisUtil) Set(key string, value string) error {
	err := r.redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisUtil) IncrBy(key string, increment int64) int64 {
	val, err := r.redisClient.IncrBy(ctx, key, increment).Result()
	if err != nil {
		panic(err)
	}
	return val
}

// HashSet 向key的hash中添加元素field的值
func (r *RedisUtil) HashSet(key, field string, data interface{}) error {
	err := r.redisClient.HSet(ctx, key, field, data).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisUtil) HashDel(key, field string) error {
	if err := r.redisClient.HDel(ctx, key, field).Err(); err != nil {
		return err
	}
	return nil
}

// BatchHashSet 批量向key的hash添加对应元素field的值
func (r *RedisUtil) BatchHashSet(key string, fields map[string]interface{}) (bool, error) {
	val, err := r.redisClient.HMSet(ctx, key, fields).Result()
	if err != nil {
		logging.Logger.Error("Redis HMSet Error:", err)
		return false, err
	}
	return val, nil
}

// HashGet 通过key获取hash的元素值
func (r *RedisUtil) HashGet(key, field string) (string, error) {
	result := ""
	val, err := r.redisClient.HGet(ctx, key, field).Result()
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
func (r *RedisUtil) BatchHashGet(key string, fields ...string) (map[string]interface{}, error) {
	resMap := make(map[string]interface{})
	for _, field := range fields {
		var result interface{}
		val, err := r.redisClient.HGet(ctx, key, fmt.Sprintf("%s", field)).Result()
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

func (r *RedisUtil) HashGetAll(key string) (map[string]string, error) {
	result, err := r.redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		logging.Logger.Error("Redis HashGetAll Error:", err)
		return nil, err
	}
	return result, nil
}

func (r *RedisUtil) PFAdd(key string, els ...interface{}) error {
	if err := r.redisClient.PFAdd(ctx, key, els...).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisUtil) PFCount(keys ...string) (int64, error) {
	result, err := r.redisClient.PFCount(ctx, keys...).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}
