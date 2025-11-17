package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/Sandwichzzy/relayer-node/cache"
	"github.com/Sandwichzzy/relayer-node/service/service"
)

// 定义 Routes 结构体，持有路由器、服务实例和缓存
type Routes struct {
	router      *chi.Mux
	svc         service.Service
	enableCache bool
	cache       *cache.LruCache
}

// NewRoutes ... Construct a new route handler instance
func NewRoutes(r *chi.Mux, svc service.Service, enableCache bool, cache *cache.LruCache) Routes {
	return Routes{
		router:      r,
		svc:         svc,
		enableCache: enableCache,
		cache:       cache,
	}
}
