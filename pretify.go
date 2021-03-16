package pretify

import (
	"encoding/json"
	"log"
)

func New(in interface{}) string {
	out, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		log.Println(err)
	}
	return string(out)
}
