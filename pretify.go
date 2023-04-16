package pretify

import (
	"bytes"
	"encoding/json"
	"log"
)

func JSON(in interface{}) string {
	if in == nil {
		return ""
	}
	out, err := json.MarshalIndent(in, "", " ")
	if err != nil {
		log.Println(err)
		return ""
	}
	return bytes.NewBuffer(out).String()
}
