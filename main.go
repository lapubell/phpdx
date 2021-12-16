package main

import "net/http"

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/template", handleTemplate)
	http.HandleFunc("/json", handleJSON)

	http.ListenAndServe(":8888", nil)
}
