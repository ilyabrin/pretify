package pretify

import (
	"encoding/json"
	"log"
)

func Pretify(input interface{}) string {
	result, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		log.Println(err)
	}
	return string(result)
}
