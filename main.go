package main

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}


func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)



}

func index(w http.ResponseWriter, r *http.Request) {

	lst := []int{}

	for x := 0; x <2; x++ {

		for i := 1; i <= 10; i++ {

			lst = append(lst, i)

		}
	}

	rand.Seed(time.Now().UnixNano())

	for i := len(lst) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		lst[i], lst[j] = lst[j], lst[i]
	}

	tpl.ExecuteTemplate(w, "index.gohtml", lst)

}
