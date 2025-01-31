package types

import (
	_ "embed"
	"encoding/json"
)

//go:embed data.json
var data string

type Region struct {
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	Children []Region `json:"children"`
}

var Regions []Region

func init() {
	if len(Regions) > 0 {
		return
	}
	if err := json.Unmarshal([]byte(data), &Regions); err != nil {
		panic(err)
	}
}
