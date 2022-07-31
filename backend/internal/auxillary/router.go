package auxiliary

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-nacelle/nacelle"
// )

// type Router struct {
// 	Logger nacelle.Logger
// 	Health nacelle.Health
// }

// type RouterAuxOption func(r *Router)

// func WithRouterLogger(logger nacelle.Logger) RouterAuxOption {
// 	return func(r *Router) {
// 		r.Logger = logger
// 	}
// }

// func WithHealth(health nacelle.Health) RouterAuxOption {
// 	return func(r *Router) {
// 		r.Health = health
// 	}
// }

// func NewRouter(opts ...RouterAuxOption) *Router {
// 	nar := Router{}

// 	for _, opt := range opts {
// 		opt(&nar)
// 	}

// 	return &nar
// }

// func (r *Router) Configure(ge gin.IRouter) {

// }
