// Package cache -----------------------------
// @file      : user.go
// @author    : hcjjj
// @contact   : hcjjj@foxmail.com
// @time      : 2024-03-18 17:04
// -------------------------------------------
package cache

import (
	"basic-go/webook/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

//var ErrKeyNotExist = errors.New("Key 不存在")

var ErrKeyNotExist = redis.Nil

type UserCache struct {
	// 传单机 Redis 可以
	// 传 cluster 的 Redis 也可以
	client     redis.Cmdable
	expiration time.Duration
}

// A 用到了 B, B 一定是接口 【面向接口编程】
// A 用到了 B, B 一定是 A 的字段 【组合】
// A 用到了 B, A 绝对不初始化 B 而是外面注入 【依赖注入】

func NewUserCache(client redis.Cmdable, expiration time.Duration) *UserCache {
	return &UserCache{
		client:     client,
		expiration: expiration,
	}
}

func (cache *UserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	// error 为 nil，即有数据
	// 如果没有数据 返回一个特定的 error
	key := cache.key(id)
	val, err := cache.client.Get(ctx, key).Bytes()
	//  redis.Nil
	if err != nil {
		return domain.User{}, err
	}
	var u domain.User
	// 反序列化
	err = json.Unmarshal(val, &u)
	return u, err
}

func (cache *UserCache) Set(ctx context.Context, u domain.User) error {
	// 序列化为 json
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	key := cache.key(u.Id)
	cache.client.Set(ctx, key, val, cache.expiration)
	return nil
}

func (cache *UserCache) key(id int64) string {
	// user:info:123
	// user_info_123
	// bumen_xiaozu_user_info_key
	return fmt.Sprintf("user:info:%d", id)
}