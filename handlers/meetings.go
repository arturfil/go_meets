package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	res := struct{
        Msg string
        Code int
    }{
		Msg:  "Health Check",
		Code: 200,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

func getAllMeetings(w http.ResponseWriter, r *http.Request) {
    
}
