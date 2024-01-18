package extract

import (
	"sp-go-service/src/Resources/global"
)

func LastError() map[string]interface{} {
	return global.LastError
}
