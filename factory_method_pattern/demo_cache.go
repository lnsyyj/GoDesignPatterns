package factory_method_pattern

// 定义一个cache的抽象接口
type Cache interface {
	Set(key, value string)
	Get(key string) string
}

// 实现具体的cache: Redis
type RedisCache struct {
	data map[string]string
}

func (this *RedisCache) Set(k, v string)  {
	this.data[k] = v
}

func (this *RedisCache) Get(k string) string {
	return this.data[k]
}

func NewRedisCache() *RedisCache {
	return &RedisCache{
		data: make(map[string]string),
	}
}

// 实现具体的cache: MemCache
type MemCache struct {
	data map[string]string
}

func (this *MemCache) Set(k, v string)  {
	this.data[k] = v
}

func (this *MemCache) Get(k string) string {
	return this.data[k]
}

func NewMemCache() *MemCache {
	return &MemCache{
		data: make(map[string]string),
	}
}

// 定义一个抽象的cache工厂
type CacheFactory interface {
	Create() Cache
}

// 实现具体的工厂: Redis
type RedisCacheFactory struct {

}

func (this *RedisCacheFactory) Create() Cache {
	return NewRedisCache()
}

// 实现具体的工厂: Mem
type MemCacheFactory struct {

}

func (this *MemCacheFactory) Create() Cache {
	return NewMemCache()
}
