package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController                 home
	registerController             register
	articlePostController		   articlepost
	//loginController		   login
	//standLocatorController standLocator
)

func Startup(templates map[string]*template.Template) {

	homeController.homeTemplate = templates["home.html"]
	registerController.registerTemplate = templates["register.html"]
	registerController.loginTemplate = templates["login.html"]
	articlePostController.postTemplate = templates["post.html"]
	articlePostController.viewTemplate = templates["view_post.html"]
	homeController.registerRoutes()
	registerController.registerRoutes()
	articlePostController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
