package martiniHandlers

import (
	"fmt"
	"github.com/codegangsta/martini"
	"io"
	"net/http"
	"time"
)

type RequestWriter struct {
	writer io.Writer
}

// RequestLogger returns a middleware handler that logs the request as it goes in and the response as it goes out.
func RequestLogger(writer io.Writer, resourceName string) martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		start := time.Now()

		rw := res.(martini.ResponseWriter)
		c.Next()

		writer.Write([]byte(fmt.Sprintf("REQUEST: %s %s%s %v in %v\n", req.Method, resourceName, req.URL.Path, rw.Status(), time.Since(start))))
	}
}
