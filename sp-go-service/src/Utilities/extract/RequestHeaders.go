package extract

import (
	"sp-go-service/src/Resources/global"
	"strings"
)

func RequestHeaders() (result string) {
	var partes []string
	for clave, valor := range global.RequestHeaders {
		partes = append(partes, clave+": "+valor)
	}
	return strings.Join(partes, ", ")
}
