package main

import (
	"html/template"
	"net/http"
)

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	t := `
<h1>This is a template! We can output variables in this with some built in tooling.</h1>
<ul>
	<li>Name: {{ .Name }}</li>
	<li>Age: {{ .Age }}</li>
	<li>Fav Languages:
		<ul>
			{{ range .Languages }}
				<li>{{ . }}</li>
			{{ end }}
		</ul>
	</li>
</ul>
`
	template, _ := template.New("html").Parse(t)
	template.Execute(w, map[string]interface{}{
		"Name":      "Kurtis",
		"Age":       37,
		"Languages": []string{"PHP", "Golang"},
	})
}
