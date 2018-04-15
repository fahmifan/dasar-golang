package main

import "fmt"
import "net/http"
import "html/template"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar")
}

func main() {
	// Defined routing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "john wick",
			"Message": "have a nice day",
		}
		// E:/Programming/Golang/projects/
		var t, err = template.ParseFiles("E:/Programming/Golang/projects/go/src/dasar-golang/48-web/template.html")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/index", index)

	PORT := ":8000"
	fmt.Print("starting web server at http://localhost", PORT, "\n")

	http.ListenAndServe(PORT, nil)
}
