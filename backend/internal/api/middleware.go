package api

// import (
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-nacelle/nacelle"
// )

// // NacelleGinLogger is a middleware handler function that inserts a nacelle
// // logger for Gin.
// func NacelleGinLogger(nl nacelle.Logger) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// before request
// 		t := time.Now()

// 		c.Next()
// 		// after request

// 		writerInfo := map[string]interface{}{
// 			"status":     c.Writer.Status(),
// 			"method":     c.Request.Method,
// 			"path":       c.Request.URL.Path,
// 			"IP":         c.ClientIP(),
// 			"latency":    time.Since(t),
// 			"user-agent": c.Request.UserAgent(),
// 		}

// 		reqStatus := c.Writer.Status()

// 		switch {
// 		case reqStatus >= 400 && reqStatus < 500:
// 			nl.WarningWithFields(writerInfo, "API-WARN")
// 		case reqStatus >= 500:
// 			nl.ErrorWithFields(writerInfo, "API-ERROR")
// 		default:
// 			nl.DebugWithFields(writerInfo, "API-DEBUG")
// 		}
// 	}
// }

// // InitHeaders is a middleware handler function that sets default headers to be
// // used by the gin routers.
// func InitHeaders() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// STS prevents Man-in-the-Middle attacks from redirecting request to a malicious website
// 		// max age is set to ensure HSTS is strictly enforcced for at least one year (in seconds)
// 		c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000")
// 		// CSP enforces that source of content (i.e. images, scripts, etc) are trusted
// 		// and allowed by specified subdomain and origin.
// 		// Header value allows content to come from the site's own origin and subdomains from humana.com
// 		c.Writer.Header().Set("Content-Security-Policy", "default-src: 'self' *.humana.com")
// 		c.Next()
// 	}
// }
