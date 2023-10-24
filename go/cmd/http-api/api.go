package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	gu2en "github.com/hirenchauhan2/gu2en/go"
)

type Message struct {
	Text string `json:"text"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		var message Message
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("%s\n", message.Text)
		text := message.Text

		result := gu2en.Transliterate(text)

		w.Write([]byte(result + "\n"))
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
