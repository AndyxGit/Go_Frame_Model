package parameter

import "sp-go-service/src/Resources/commons"

var HTTP_PORT = commons.SearchEnv("HTTP_PORT", "8080")
var LOG_LEVEL = commons.SearchEnv("LOGLEVEL", "DEBUG")
var DATE_FORMAT = commons.SearchEnv("DATE_FORMAT", "2006-01-02 15:04:05 -0700")
var LINE_RECHARGE_ENDPOINT_PATH = commons.SearchEnv("LINE_RECHARGE_ENDPOINT_PATH", "/lineRecharge")
var RESPONSE_TIMEOUT_SEC = commons.SearchEnv("RESPONSE_TIMEOUT_SEC", "30")

var DBCP_CCARD_DATACONSUMPTIONSERVICE_DRIVER = commons.SearchEnv("DBCP_CCARD_DATACONSUMPTIONSERVICE_DRIVER", "oracle.jdbc.driver.OracleDriver")
var DBCP_CCARD_DATACONSUMPTIONSERVICE_URL = commons.SearchEnv("DBCP_CCARD_DATACONSUMPTIONSERVICE_URL", "bengolea.claro.amx:1521:DCCARD")
var DBCP_CCARD_DATACONSUMPTIONSERVICE_USERNAME = commons.SearchEnv("DBCP_CCARD_DATACONSUMPTIONSERVICE_USERNAME", "dataconsumptionservice")
var DBCP_CCARD_DATACONSUMPTIONSERVICE_PASSWORD = commons.SearchEnv("DBCP_CCARD_DATACONSUMPTIONSERVICE_PASSWORD", "dataconsumptionservice")
