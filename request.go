package martiniHandlers

import (
	"github.com/codegangsta/martini"
	"github.com/coreos/go-log/log"
	"net/http"
	"time"
)

// RequestLogger returns a middleware handler that logs the request as it goes in and the response as it goes out.
func RequestLogger() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context, logger *log.Logger) {
		start := time.Now()

		rw := res.(martini.ResponseWriter)
		c.Next()

		logger.Infof("[REQ] %v %s %s in %v", rw.Status(), req.Method, req.URL.Path, time.Since(start))
	}
}
