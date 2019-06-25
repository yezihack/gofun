package tools

import (
	"errors"
	"sync"
	"time"
)

const (
	DefaultExpiration time.Duration = 0
)

type item struct {
	Expired int64       //过期时间
	Value   interface{} //存储的值
}

//判断是否过期
func (m item) IsExpired() bool {
	if m.Expired == 0 {
		return false
	}
	return time.Now().UnixNano() > m.Expired
}

type Cache struct {
	*goCache
}

type goCache struct {
	DefaultExpiration time.Duration
	items             map[string]item //key => Item
	lock              sync.RWMutex
}

//实例对象
func New(defaultTTL time.Duration) *Cache {
	m := make(map[string]item)
	c := &goCache{
		DefaultExpiration: defaultTTL,
		items:             m,
	}
	cc := &Cache{
		goCache: c,
	}
	return cc
}

//缓存某值
func (c *Cache) Put(key string, value interface{}, ttl time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()
	var ex int64
	if ttl == DefaultExpiration {
		ttl = c.DefaultExpiration
	}
	if ttl > 0 {
		ex = time.Now().Add(ttl).UnixNano()
	}
	item := item{
		Expired: ex,
		Value:   value,
	}
	c.items[key] = item
}

//默认操作
func (c *Cache) PutDefault(key string, value interface{}) {
	c.Put(key, value, DefaultExpiration)
}

//获取某键值
func (c *Cache) Get(key string) (reply interface{}, err error) {
	item, isExist := c.items[key]
	if !isExist {
		err = errors.New("key is not exists")
		return
	}
	if item.Expired > 0 {
		if item.IsExpired() {
			err = errors.New("key is expired")
			return
		}
	}
	reply = item.Value
	return
}

//删除某键值
func (c *Cache) Delete(key string) bool {
	delete(c.items, key)
	return true
}

//判断键是否存在
func (c *Cache) Has(key string) bool {
	_, isExist := c.items[key]
	return isExist
}
