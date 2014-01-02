package martiniReturnHandlers

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

func JsonReturnHandler(logger *log.Logger) martini.ReturnHandler {
	return func(res http.ResponseWriter, vals []reflect.Value) {

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

			if data, err = json.Marshal(responseVal.Interface()); err != nil {
				data = []byte("INTERNAL ERROR")
				statusCode = http.StatusInternalServerError
				logger.Printf("Error marshalling json: %v", err)
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
