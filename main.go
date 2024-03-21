package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

var personList []Person

func main() {
	data, _ := os.ReadFile("./data.json")
	json.Unmarshal(data, &personList)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./index.html")
		tmpl.Execute(w, personList)
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is Listening on 8080...")
}
