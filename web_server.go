package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, API!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "john wick",
			"Message": "have a nice day",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8888/")
	http.ListenAndServe(":8888", nil)

}
