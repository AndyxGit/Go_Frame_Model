package attach

import (
	"sp-go-service/src/Resources/global"
)

func RequestHeaders(fieldName string, fieldValue string) {
	global.RequestHeaders[fieldName] = fieldValue
}
