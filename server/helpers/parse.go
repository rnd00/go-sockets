package helpers

// helpers - parse.go
// - run with handleConnection
// - obj: determine whether usable json or not
// - obj: handling json stuffs from received bytes

import (
	"encoding/json"
	"log"
)

// FnCall is a model which we're basing the data into
type FnCall struct {
	Function string        `json:"function"`
	Data     []interface{} `json:"data"`
}

// Data here will be used for FnCall struct if running on 2-step
type Data struct {
	Data []interface{} `json:"data"`
}

// ParseFnCall is a json.Unmarshal wrapper
func ParseFnCall(data []byte) {
	var r FnCall
	err := json.Unmarshal(data, &r)
	if err != nil {
		log.Println(err)
	}
	// log.Println(r.Function)
	// log.Println(r.Data)
	FunctionCaller(r)
}

// FunctionCaller will reroute the function calls elsewhere
func FunctionCaller(data FnCall) {
	fName := data.Function
	fData := data.Data
	switch fName {
	case "HW":
		HelloWorld()
	case "H":
		Hello(fData)
	default:
		log.Printf("Unrecognized function (%s) was called!", fName)
	}
}

// ParseData is a json.Unmarshal wrapper for 'Data'
func ParseData(data []byte) Data {
	var n Data
	err := json.Unmarshal(data, &n)
	if err != nil {
		log.Println(err)
	}

	return n
}
