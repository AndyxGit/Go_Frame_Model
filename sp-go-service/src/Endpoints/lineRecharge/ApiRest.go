package lineRecharge

import (
	"fmt"
	"net/http"
	"sp-go-service/src/Resources/global"
	"sp-go-service/src/Resources/logger"
	"sp-go-service/src/Resources/structure"
	"sp-go-service/src/Utilities/tool"
	"sp-go-service/src/Utilities/verify"
	"strings"
	"time"
)

/*
// RequestData representa los datos que esperas en la solicitud
type RequestData struct {
	Param1 string `json:"param1"`
	Param2 int    `json:"param2"`
}
*/

func ApiRest(method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tool.CleanGlobalEnvs()

		if !verify.StructHeaders(r, structure.RequestHeaderStruct{}) {
			tool.ResponseHttp(w)
			return
		} else {
			logger.Console("DEBUG", "", "Estructura Headers validada correctamente")
		}

		logger.Console("DEBUG", "", "Inicia validacion de Method permitido")
		if r.Method != method {
			time.Sleep(10 * time.Second)
			fmt.Printf("\n --------------------------------------------- \n")
			fmt.Println(global.RequestHeaders)
			fmt.Println(global.ResponseBody)
			fmt.Println(global.LastError)
			fmt.Printf("--------------------------------------------- \n\n")

			logger.Console("ERROR", strings.ToUpper(fmt.Sprintf("ERROR_BADREQUEST_%s_NOT_VALID_METHOD", r.Method)), r.URL.Path)
			tool.ResponseHttp(w)
			return
		} else {
			logger.Console("DEBUG", "", fmt.Sprintf("Method %v Permitido", r.Method))
		}

		/*
			var requestData RequestData
			if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		*/

		Service(w)

	}
}
