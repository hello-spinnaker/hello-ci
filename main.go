package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	colour := struct {
		BackgroundColour string
		Colour           string
		Name             string
	}{BackgroundColour: "#444444", Colour: "#888888", Name: "grey"}

	t, err := template.ParseFiles("page.tpl")

	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			t.Execute(w, colour)
		},
	)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("BLANKPAGE_PORT")), nil)
}
