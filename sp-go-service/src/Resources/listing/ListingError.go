package listing

import "sp-go-service/src/Resources/structure"

var ListingError = []structure.ListingErrorStruct{

	{Id: "ERROR_FUNCTION_STRUCTOBJECT_PARAMNOTSTRUCT", Code: 900400, Reason: "Explanation: the function requires a slice structure to work"},

	// Metodo no permitido
	{Id: "ERROR_BADREQUEST_GET_NOT_VALID_METHOD", Code: 100501, Reason: "Explanation: Method GET Not Valid"},
	{Id: "ERROR_BADREQUEST_POST_NOT_VALID_METHOD", Code: 100502, Reason: "Explanation: Method POST Not Valid"},
	{Id: "ERROR_BADREQUEST_PUT_NOT_VALID_METHOD", Code: 100503, Reason: "Explanation: Method PUT Not Valid"},
	{Id: "ERROR_BADREQUEST_DELETE_NOT_VALID_METHOD", Code: 100504, Reason: "Explanation: Method DELETE Not Valid"},
	{Id: "ERROR_BADREQUEST_HEAD_NOT_VALID_METHOD", Code: 100505, Reason: "Explanation: Method HEAD Not Valid"},

	// BadRequest Grupo de error 1004XX
	{Id: "ERROR_BADREQUEST_SESSIONID_MISSING", Code: 100401, Reason: "Explanation: Missing Header"},
	{Id: "ERROR_BADREQUEST_CHANNELID_MISSING", Code: 100402, Reason: "Explanation: Missing Header"},
	{Id: "ERROR_BADREQUEST_USERID_MISSING", Code: 100403, Reason: "Explanation: Missing Header"},
	{Id: "ERROR_BADREQUEST_NUMBER_MISSING", Code: 100404, Reason: "Explanation: Missing Header"},
	{Id: "ERROR_BADREQUEST_DATA_MISSING", Code: 100405, Reason: "Explanation: Missing Header"},

	//Header de tipo no valido - Grupo de error 1005XX (Los Headers String siempre son string asi que no hace falta valida tipo)
	//Los Errores pueden ser (NOTINT, NOTBOOL, NOTFLOAT, NOTFORMATDATE) segun corresponda
	{Id: "ERROR_BADREQUEST_NUMBER_NOTINT", Code: 100500, Reason: "Explanation: Missing Header"},
	{Id: "ERROR_BADREQUEST_DATA_NOTFLOAT", Code: 100405, Reason: "Explanation: Missing Header"},

	// Errores Desconocido no tipificado
	{Id: "ERROR_UNKNOWN", Code: 900999, Reason: "Explanation: Constant not found in ListingErrorStruct"},

	//errores de prueba
	{Id: "ERROR_BLA_BLA", Code: 100100, Reason: "aquí se dice la razón"},
	{Id: "ERROR_XYZ", Code: 100100, Reason: "aquí se dice la razón"},
	{Id: "ERROR_XYZ", Code: 100100, Reason: "aquí se dice la razón"},
}
