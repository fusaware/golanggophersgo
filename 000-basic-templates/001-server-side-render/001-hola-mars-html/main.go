package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html") // because we are diligent

		d := struct {
			FlatSurface string
		}{
			"Mars",
		}

		t, err := template.ParseFiles("hola-mars.gohtml")
		if err != nil {
			panic(fmt.Sprintf("doh: %v", err))
		}
		if err = t.Execute(w, d); err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/ionlyspeakenglish", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html") // because we are diligent

		d := struct {
			FlatSurface string
		}{
			"Mundo",
		}

		t, err := template.ParseFiles("hello-world.gohtml")
		if err != nil {
			panic(fmt.Sprintf("doh: %v", err))
		}

		if err = t.Execute(w, d); err != nil {
			log.Fatal(err)
		}
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("don't like when this happens: %v", err)
	}
}
