package verify

import (
	"fmt"
	"net/http"
	"reflect"
	"sp-go-service/src/Resources/logger"
	"sp-go-service/src/Resources/parameter"
	"sp-go-service/src/Utilities/attach"
	"strconv"
	"strings"
	"time"
)

func StructHeaders(r *http.Request, StructCheck interface{}) (isOK bool) {

	logger.Console("DEBUG", "", "Se inicia validacion de objeto StrucCheck")
	val := reflect.ValueOf(StructCheck)
	if val.Kind() != reflect.Struct {
		logger.Console("ERROR", "ERROR_FUNCTION_STRUCTOBJECT_PARAMNOTSTRUCT", "")
		return false
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		headerValue := r.Header.Get(field.Name)

		logger.Console("DEBUG", "", fmt.Sprintf("Se inicia validacion de existencia del Header %v", field.Name))
		if headerValue == "" {
			logger.Console("ERROR", strings.ToUpper(fmt.Sprintf("ERROR_BADREQUEST_%s_MISSING", field.Name)), field.Name)
			return false
		}

		logger.Console("DEBUG", "", fmt.Sprintf("Se inicia validacion del tipo de dato del Header %v", field.Name))
		switch field.Type.Kind() {
		case reflect.String:

		case reflect.Int, reflect.Int64:
			if _, err := strconv.ParseInt(headerValue, 10, 64); err != nil {
				logger.Console("ERROR", strings.ToUpper(fmt.Sprintf("ERROR_BADREQUEST_%s_NOTINT", field.Name)), field.Name)
				return false
			}

		case reflect.Bool:
			if _, err := strconv.ParseBool(headerValue); err != nil {
				logger.Console("ERROR", strings.ToUpper(fmt.Sprintf("ERROR_BADREQUEST_%s_NOTBOOL", field.Name)), field.Name)
				return false
			}

		case reflect.Float32, reflect.Float64:
			if _, err := strconv.ParseFloat(headerValue, 64); err != nil {
				logger.Console("ERROR", strings.ToUpper(fmt.Sprintf("ERROR_BADREQUEST_%s_NOTFLOAT", field.Name)), field.Name)
				return false
			}

		case reflect.Struct:
			if field.Type == reflect.TypeOf(time.Time{}) {
				if _, err := time.Parse(parameter.DATE_FORMAT, headerValue); err != nil {
					logger.Console("ERROR", strings.ToUpper(fmt.Sprintf("ERROR_BADREQUEST_%s_NOTFORMATDATE", field.Name)), field.Name)
					return false
				}
			}

		}

		attach.RequestHeaders(field.Name, headerValue)
	}
	return true
}
