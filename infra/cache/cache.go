package cache

import (
	"errors"
	"github.com/mediocregopher/radix/v3"
)

// Serializer interface.
type Serializer interface {
	EncodeToString(v interface{}) (string, error)
	DecodeFromString(s string, v interface{}) error
}

// Redis cache implementation.
type RedisCache struct {
	p          *radix.Pool
	serializer Serializer
}

func NewCacheRedis(network string, addr string, size int, serializer Serializer) *RedisCache {
	p, err := radix.NewPool(network, addr, size)
	if err != nil {
		panic(err.Error())
	}

	r := &RedisCache{
		p:          p,
		serializer: serializer,
	}

	return r
}

func (r *RedisCache) Get(key string, obj any) error {
	var result string
	err := r.p.Do(radix.Cmd(&result, "GET", key))
	if err != nil {
		return err
	}

	if result == "" {
		return errors.New("cache key " + key + " not found")
	}

	err = r.serializer.DecodeFromString(result, obj)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCache) Set(key string, value any) error {
	serialized, err := r.serializer.EncodeToString(value)

	if err != nil {
		return err
	}
	err = r.p.Do(radix.Cmd(nil, "SET", key, serialized))

	if err != nil {
		return err
	}

	return nil
}
