package factory_method_pattern

import "testing"

func TestRedisCache(t *testing.T) {

	var redisCacheFactory CacheFactory
	redisCacheFactory = &RedisCacheFactory{}
	redisCache := redisCacheFactory.Create()

	redisCache.Set("key1", "value1")
	value := redisCache.Get("key1")

	t.Log(value)
}