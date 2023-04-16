package pretify

import (
	"bytes"
	"encoding/json"
	"log"
)

func JSON(input interface{}) string {
	if input == nil {
		return ""
	}

	out, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		log.Println(err)
		return ""

	}

	return bytes.NewBuffer(out).String()
}
