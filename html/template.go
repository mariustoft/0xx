package main

import (
    "html/template"
    "net/http"
)

type PageData struct {
    Title string
    Body  string
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("layout.html"))
        data := PageData{
            Title: "Hello, World!",
            Body:  "Welcome to our website!",
        }
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":8080", nil)
}