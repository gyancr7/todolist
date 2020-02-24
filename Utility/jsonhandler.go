package Utility

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UnmarshalInput(jsonData []byte, entity interface{}) {
	err := json.Unmarshal([]byte(string(jsonData)), &entity)

	if err != nil {
		fmt.Printf("Error Unmarshaling")
	}
}

func MarshalOutput(entity interface{}, response http.ResponseWriter) {
	err := json.NewEncoder(response).Encode(&entity)

	if err != nil {
		fmt.Printf("Error Marshaling")
	}
}