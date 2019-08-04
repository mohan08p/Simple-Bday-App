package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Custom catch method...
func Catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// RespondwithError return error message
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondwithJSON(w, code, map[string]string{"message": msg})
}

// RespondwithJSON write json response format
func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
