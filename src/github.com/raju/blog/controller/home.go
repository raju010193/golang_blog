package controller


import (
	_ "fmt"
	"html/template"
	_ "log"
	"net/http"
	"github.com/raju/blog/viewmodel"
)

type home struct {
	homeTemplate  *template.Template
	//loginTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	//http.HandleFunc("/login", h.handleLogin)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	w.Header().Add("Content-Type", "text/html")
	h.homeTemplate.Execute(w, vm)
}


