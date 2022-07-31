package api

// import (
// 	auxiliary "interview/internal/auxillary"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-nacelle/nacelle"
// )

// // Factory is a structure used to create and launch multiple routers.
// // These routers are configured from the middleware and are passed in services
// // to use when responding to their requests.
// type Factory struct {
// 	logger nacelle.Logger
// 	health nacelle.Health
// }

// // NewFactory creates a new default value Factory that accepts extra factory
// // configuration functions.
// func NewFactory(logger nacelle.Logger, opts ...FactoryOption) *Factory {
// 	napi := Factory{}
// 	napi.logger = logger

// 	for _, opt := range opts {
// 		opt(&napi)
// 	}

// 	return &napi
// }

// // FactoryOption represents a function that is passed into the NewFactory
// // function to configure the router factory.
// type FactoryOption func(api *Factory)

// // WithHealth passes in the default nacelle health.
// func WithHealth(health nacelle.Health) FactoryOption {
// 	return func(api *Factory) {
// 		api.health = health
// 	}
// }

// // ConfigureRouters creates the individual routers.
// func (a *Factory) ConfigureRouters(ge gin.IRouter) {

// 	// Auxiliary routes
// 	auxRouter := auxiliary.NewRouter(
// 		auxiliary.WithRouterLogger(a.logger),
// 		auxiliary.WithHealth(a.health),
// 	)
// 	auxRouter.Configure(ge.Group("/"))
// }
