package utils

import (
	"encoding/json"
)

func MapStruct(dest interface{}, src interface{}) {
	data, _ := json.Marshal(src)
	_ = json.Unmarshal(data, dest)
}
