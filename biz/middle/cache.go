package middle

import (
	"github.com/coocood/freecache"
)

var Cache *freecache.Cache

// 缓存优化
func init() {
	Cache = freecache.NewCache(20 * 1024 * 1024)
}

func AddCode(user, code string) {
	_ = Cache.Set([]byte(user), []byte(code), 60)
}
