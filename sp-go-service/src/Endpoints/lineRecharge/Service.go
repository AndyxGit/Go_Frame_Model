package lineRecharge

import (
	"net/http"
	"sp-go-service/src/Resources/global"
	"sp-go-service/src/Resources/logger"
	"sp-go-service/src/Utilities/tool"
)

func Service(w http.ResponseWriter) {

	logger.Console("DEBUG", "", "Inicia Logica Interna del Servicio")

	global.ResponseBody["FieldA"] = "ValueA"
	global.ResponseBody["FieldB"] = "ValueB"
	global.ResponseBody["FieldC"] = "ValueC"

	tool.ResponseHttp(w)
	logger.Console("DEBUG", "", "Termina Logica Interna del Servicio")
}
