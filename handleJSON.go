package main

import (
	"encoding/json"
	"net/http"
)

func handleJSON(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Thing              string `json:"thing"`
		AnswerToEverything int
		hiddenFromJSON     bool // this is unexported!
	}

	stuff := data{
		Thing:              "I'm some stuff!",
		AnswerToEverything: 42,
		hiddenFromJSON:     true,
	}

	bytes, _ := json.Marshal(stuff)

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
