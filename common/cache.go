package common

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache *cache.Cache

/**
初始化缓存对象 5分钟 有效 每10分钟清理一下缓存
 */
func init() {
	Cache = cache.New(5*time.Minute, 10*time.Minute)
}
