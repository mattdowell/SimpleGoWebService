package json

import (
	"GoWebService/src/basicdb/obs"
	"encoding/json"
)

/*
   Converts the struct to a JSON string
 */
func Marshall(inObj obs.SimpleDbType) string {
	emp, _ := json.Marshal(inObj)
	theReturn := string(emp)
	return theReturn
}

/**
  Converts a json tring to a type
 */
func UnMarshall(inJsonString string) obs.SimpleDbType {
	bytes := []byte(inJsonString)
	var res obs.SimpleDbType
	json.Unmarshal(bytes, &res)
	return res
}
