package sqldb

import (
	"database/sql"
	"fmt"
	"sp-go-service/src/Resources/parameter"
	"strings"
)

func GetQuery() {

	// Obtener variables de entorno
	url := parameter.DBCP_CCARD_DATACONSUMPTIONSERVICE_URL
	username := parameter.DBCP_CCARD_DATACONSUMPTIONSERVICE_USERNAME
	password := parameter.DBCP_CCARD_DATACONSUMPTIONSERVICE_PASSWORD

	host := strings.Split(url, ":")[0]
	port := strings.Split(url, ":")[1]
	base := strings.Split(url, ":")[2]
	dsn := fmt.Sprintf("%s/%s@%s:%s/%s", username, password, host, port, base)

	//fmt.Println(host, port, base, dsn)

	// Abrir conexión a la base de datos OJO VAS POR AQUI
	db, err := sql.Open("godror", dsn)
	if err != nil {
		fmt.Println(db)

		//logger.Console("ERROR", fmt.Sprintf("Error al abrir la base de datos: %v", err), "psp000000000000", "PSP", "UserLocal")
	}
	defer db.Close()

}
