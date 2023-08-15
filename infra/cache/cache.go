package cache

import (
	"errors"
	"github.com/mediocregopher/radix/v3"
)

// Redis cache implementation.
type RedisCache struct {
	p *radix.Pool
}

func NewCacheRedis(network string, addr string, size int) *RedisCache {
	p, err := radix.NewPool(network, addr, size)
	if err != nil {
		panic(err.Error())
	}

	r := &RedisCache{
		p: p,
	}

	return r
}

func (r *RedisCache) Get(key string) (string, error) {
	var result string
	err := r.p.Do(radix.Cmd(&result, "GET", key))
	if err != nil {
		return "", err
	}

	if result == "" {
		return "", errors.New("cache key " + key + " not found")
	}

	return result, nil
}

func (r *RedisCache) Set(key string, value string) error {
	err := r.p.Do(radix.Cmd(nil, "SET", key, value))

	if err != nil {
		return err
	}

	return nil
}