package controller


import (
	_ "fmt"
	"html/template"
	_ "log"
	"net/http"
	"github.com/raju/blog/viewmodel"
)

type register struct {
	registerTemplate  *template.Template
	loginTemplate *template.Template
}

func (h register) registerRoutes() {
	http.HandleFunc("/register", h.handleRegister)
	http.HandleFunc("/login", h.handleLogin)
}

func (h register) handleRegister(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewRegister()
	w.Header().Add("Content-Type", "text/html")
	h.registerTemplate.Execute(w, vm)
}


func (h register)handleLogin(w http.ResponseWriter, r *http.Request){
	vm := viewmodel.NewLogin()
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w,vm)
}