package launch

import (
	"fmt"
	"log"
	"net/http"
	"sp-go-service/src/Endpoints/lineRecharge"
	"sp-go-service/src/Resources/logger"
	"sp-go-service/src/Resources/parameter"
)

func Application() {
	http.HandleFunc(parameter.LINE_RECHARGE_ENDPOINT_PATH, lineRecharge.ApiRest("POST"))
	logger.Console("INFO", "", fmt.Sprintf("Servidor iniciado en el puerto TCP:%v", parameter.HTTP_PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", parameter.HTTP_PORT), nil))
}
