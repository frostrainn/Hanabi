package Utils

import (
	"encoding/json"
	"fmt"
)

//结构体转JSON
func StructToJSON(p interface{}) string {
	data, err := json.Marshal(p)

	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return "structToJSON ERROR !!!"
	}

	return string(data)
}
