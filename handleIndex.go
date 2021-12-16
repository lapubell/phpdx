package main

import "net/http"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	m := NewMascot("lil gopher buddy")

	// html inside of a string is kinda gross. Take a look at handleTemple to
	// see how we can use the built in html.Template package to do this better.
	w.Write([]byte(
		"<h1>This is the home page!</h1>" +
			"<p>Our mascot is named '" + m.name + "'</p>",
	))
}
