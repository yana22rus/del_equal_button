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


type data struct {

	Id  []int
	Name []int
}

func index(w http.ResponseWriter, r *http.Request) {

	name := []int{}

	id := []int{}

	for x := 0; x <2; x++ {

		for i := 1; i <= 10; i++ {

			name = append(name, i)

		}
	}

	rand.Seed(time.Now().UnixNano())

	for i := len(name) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		name[i], name[j] = name[j], name[i]
	}

	for i :=1; i<=20;i++{

		id = append(id,i)

	}

	render := data{id,name}

	render_data := []data{render}

	tpl.ExecuteTemplate(w, "index.gohtml", render_data)

}