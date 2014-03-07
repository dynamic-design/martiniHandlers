package martiniHandlers

import (
	"github.com/codegangsta/inject"
	"github.com/codegangsta/martini"

	"encoding/json"
	// "fmt"
	"net/http"
	"reflect"
	"strconv"
)

func JsonReturnHandler() martini.ReturnHandler {
	return func(ctx martini.Context, vals []reflect.Value) {
		rv := ctx.Get(inject.InterfaceOf((*http.ResponseWriter)(nil)))
		res := rv.Interface().(http.ResponseWriter)

		// Default status code to error to not fool client with unsuccessfull 200 requests
		var statusCode int = http.StatusInternalServerError

		// Response should always be json encoded
		res.Header().Set("Content-Type", "application/json")

		// Check if handler returned a statuscode with the response
		var responseVal reflect.Value
		if len(vals) > 1 && vals[0].Kind() == reflect.Int {
			statusCode = int(vals[0].Int())
			responseVal = vals[1]
		} else if len(vals) > 0 {
			responseVal = vals[0]
		}

		// Response should be a generic interface
		if responseVal.Kind() == reflect.Interface {
			var data []byte
			var err error

			if !responseVal.IsNil() {
				if data, err = json.Marshal(responseVal.Interface()); err != nil {
					data = []byte("INTERNAL ERROR")
					statusCode = http.StatusInternalServerError
				}
			}

			res.Header().Set("Content-Length", strconv.Itoa(len(data)))
			res.WriteHeader(statusCode)
			res.Write(data)
			return
		}
		if isByteSlice(responseVal) {
			res.Write(responseVal.Bytes())
		} else {
			res.Write([]byte(responseVal.String()))
		}
	}
}

func isByteSlice(val reflect.Value) bool {
	return val.Kind() == reflect.Slice && val.Type().Elem().Kind() == reflect.Uint8
}
