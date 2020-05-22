package controller


import (
	_ "fmt"
	"html/template"
	_ "log"
	"net/http"
	"github.com/raju/blog/viewmodel"
)

type articlepost struct {
	postTemplate  *template.Template
	viewTemplate *template.Template
}

func (h articlepost) registerRoutes() {
	http.HandleFunc("/post", h.handlePost)
	http.HandleFunc("/viewpost", h.handleViewPost)
}

func (h articlepost) handlePost(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewPost()
	w.Header().Add("Content-Type", "text/html")
	h.postTemplate.Execute(w, vm)
}


func (h articlepost)handleViewPost(w http.ResponseWriter, r *http.Request){
	vm := viewmodel.ViewPost()
	w.Header().Add("Content-Type", "text/html")
	h.viewTemplate.Execute(w,vm)
}