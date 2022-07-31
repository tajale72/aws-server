package api

// import (
// 	"context"
// 	"crypto/tls"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"strings"
// 	"sync"
// 	"syscall"
// 	"time"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-contrib/pprof"
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-nacelle/nacelle"
// 	"github.com/opentracing-contrib/go-gin/ginhttp"
// 	"github.com/opentracing/opentracing-go"
// 	ginprometheus "github.com/zsais/go-gin-prometheus"
// )

// type APIConfig struct {
// 	Port     string `env:"WEB_PORT" default:"9000"`
// 	RootPath string `env:"WEB_ROOT_PATH" default:"/gatekeeper"`

// 	TLSCert []byte `env:"TLS_CERT"`
// 	TLSKey  []byte `env:"TLS_KEY"`
// }
// type ConfigCORS struct {
// 	AllowOrigins     []string      `env:"ALLOW_ORIGINS" default:"[\"http://localhost:9000\", \"http://localhost\"]"`
// 	AllowMethods     []string      `env:"ALLOW_METHODS" default:"[\"*\"]"`
// 	AllowHeaders     []string      `env:"ALLOW_HEADERS" default:"[\"*\"]"`
// 	ExposeHeaders    []string      `env:"EXPOSE_HEADERS"`
// 	AllowCredentials bool          `env:"ALLOW_CREDENTIALS"`
// 	MaxAge           time.Duration `env:"MAX_AGE"`
// }

// var (
// 	_ nacelle.Process     = &Process{}
// 	_ nacelle.Initializer = &Process{}
// )

// // Process is used as a way to spin up routers for Gatekeeper.
// // Services from cmd/gatekeeper are passed down into this Process as tagged
// // fields.
// type Process struct {
// 	// Default health and logger for nacelle.
// 	Logger nacelle.Logger `service:"logger"`
// 	Health nacelle.Health `service:"health"`
// 	// Custom added services that are passed down.

// 	// Configuration structures.
// 	apiCfg  *APIConfig
// 	corsCfg cors.Config
// 	// http server used for the routers.
// 	APIServer *http.Server
// 	// Ways to kill and stop the Process.
// 	stopChan chan struct{}
// 	stopOnce sync.Once
// 	killChan chan os.Signal
// }

// // NewProcess creates and initializes an empty Process.
// func NewProcess() *Process {
// 	return &Process{
// 		Logger:    nacelle.NewNilLogger(),
// 		stopChan:  make(chan struct{}),
// 		killChan:  make(chan os.Signal, 1),
// 		APIServer: &http.Server{},
// 	}
// }

// // Init is ran as an initialization function for this Process.
// // Configurations are loaded and the http server is created.
// func (p *Process) Init(config nacelle.Config) error {
// 	// Load the APIFactory router config
// 	apiCfg := &APIConfig{}
// 	if err := config.Load(apiCfg); err != nil {
// 		return err
// 	}
// 	p.apiCfg = apiCfg

// 	// HTTPS server configuration currently set to skip over if not provided.
// 	var tlsCfg *tls.Config
// 	if len(p.apiCfg.TLSCert) > 0 && len(p.apiCfg.TLSKey) > 0 {
// 		// Creates a certificate from PEM encoded data.
// 		cert, err := tls.X509KeyPair(p.apiCfg.TLSCert, p.apiCfg.TLSKey)
// 		if err != nil {
// 			return fmt.Errorf("could not load tls pair: %s", err)
// 		}
// 		tlsCfg = &tls.Config{
// 			Certificates: []tls.Certificate{cert},
// 		}
// 		p.APIServer.TLSConfig = tlsCfg
// 	}
// 	p.APIServer.Addr = fmt.Sprintf(":%s", p.apiCfg.Port)

// 	// Load the CORS config options
// 	corsCfg := &ConfigCORS{}
// 	if err := config.Load(corsCfg); err != nil {
// 		return err
// 	}
// 	p.corsCfg = cors.Config{
// 		AllowOrigins:     corsCfg.AllowOrigins,
// 		AllowMethods:     corsCfg.AllowMethods,
// 		AllowHeaders:     corsCfg.AllowHeaders,
// 		ExposeHeaders:    corsCfg.ExposeHeaders,
// 		AllowCredentials: corsCfg.AllowCredentials,
// 		MaxAge:           corsCfg.MaxAge,
// 	}
// 	// Set the kill channel to be notified when signals are read.
// 	signal.Notify(p.killChan, syscall.SIGINT, syscall.SIGTERM)

// 	return nil
// }

// // InitRouter initializes the gin routers with headers, loggers, configs, etc.
// // Also this function is used to initialize prometheus/metrics.
// func (p *Process) InitRouter() *gin.Engine {
// 	gin.SetMode(gin.ReleaseMode)
// 	r := gin.New()
// 	gin.DefaultWriter = ioutil.Discard

// 	// Use CORS
// 	r.Use(cors.New(p.corsCfg))

// 	// Use custom headers
// 	r.Use(InitHeaders())

// 	// Use the MS framework logger
// 	r.Use(NacelleGinLogger(p.Logger))

// 	// Jaeger Tracing
// 	tracer := opentracing.GlobalTracer()
// 	r.Use(ginhttp.Middleware(tracer))

// 	// PPROF debugging
// 	pprof.Register(r, p.apiCfg.RootPath+"/debug/pprof")

// 	// Prometheus
// 	hostname, err := os.Hostname()
// 	if err != nil {
// 		p.Logger.Warning("get hostname:", err)
// 	}
// 	// The prefix must meet the convention where there are no hyphens
// 	promPrefix := strings.ReplaceAll(hostname, "-", "_")
// 	prom := ginprometheus.NewPrometheus(promPrefix)
// 	prom.MetricsPath = p.apiCfg.RootPath + "/metrics"

// 	// prom.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
// 	// 	url := c.Request.URL.Path
// 	// 	for _, p := range c.Params {
// 	// 		// Range over url metric replacers
// 	// 		// k, v := range controller.URLReplacers {
// 	// 		//	url = strings.Replace(url, p.Value, k, 1)
// 	// 		//	break
// 	// 		//}
// 	// 	}
// 	// 	return url
// 	// }
// 	prom.Use(r)

// 	// ConfigureRoutes the APIFactory implementation
// 	apiImpl := NewFactory(
// 		p.Logger,
// 	)

// 	// Add the versioned business logic routes
// 	apiImpl.ConfigureRouters(r.Group(p.apiCfg.RootPath))

// 	return r
// }

// // Start is called when nacelle boots and is used to begin the routers and
// // listen/serve to API requests.
// func (p *Process) Start() error {
// 	// Initialize the routers.
// 	p.APIServer.Handler = p.InitRouter()
// 	// Listen and serve using the http server.
// 	go func() {
// 		if err := p.APIServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			p.Logger.Error("http API server listen error:", err)
// 		}
// 	}()
// 	// Wait for a response from any of the channels.
// 	select {
// 	case <-p.killChan:
// 		p.Logger.Info("http API server received kill signal...")
// 		if err := p.serverShutdown(); err != nil {
// 			p.Logger.Error("http API server error shutting down:", err)
// 		}
// 	case <-p.stopChan:
// 		p.Logger.Info("http API server received stop signal...")
// 		if err := p.serverShutdown(); err != nil {
// 			p.Logger.Error("http API server error shutting down:", err)
// 		}
// 	}

// 	return nil
// }

// func (p *Process) serverShutdown() error {
// 	p.Logger.Info("http API server shutdown sequence triggered...")

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	// Attempt to shutdown the http server with a timeout period.
// 	if err := p.APIServer.Shutdown(ctx); err != nil {
// 		return err
// 	}
// 	p.Logger.Info("...http API server exit complete.")

// 	return nil
// }

// // Stop is called when nacelle begins its shutdown process.
// // This usually happens when a kill signal is read.
// func (p *Process) Stop() error {
// 	// Close the Process.
// 	p.stopOnce.Do(func() {
// 		close(p.stopChan)
// 	})
// 	return nil
// }
