package main

import (
	"./app"
	"./controllers"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
)

func htmlIndexPageTemplateHandler(w http.ResponseWriter, r *http.Request) {
	htmlTemplate, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	htmlTemplate.Execute(w, nil)
}


func htmlRegistrationPageTemplateHandler(w http.ResponseWriter, r *http.Request) {
	htmlTemplate, err := template.ParseFiles("templates/registration.html")

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	htmlTemplate.Execute(w, nil)
}

func htmlVisitPageTemplateHandler(w http.ResponseWriter, r *http.Request) {
	htmlTemplate, err := template.ParseFiles("templates/visit.html")

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	htmlTemplate.Execute(w, nil)
}

func html404PageTemplateHandler(w http.ResponseWriter, r *http.Request) {
	htmlTemplate, err := template.ParseFiles("templates/404.html")

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	htmlTemplate.Execute(w, nil)
}


func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", htmlIndexPageTemplateHandler)
	router.HandleFunc("/registration", htmlRegistrationPageTemplateHandler)
	router.HandleFunc("/visit", htmlVisitPageTemplateHandler)
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/actions/new", controllers.CreateAction).Methods("POST")
	router.HandleFunc("/api/me/actions", controllers.GetActionsFor).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(html404PageTemplateHandler)

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "60000"
	}

	fmt.Println("\nNow to see the application open http://localhost:",port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
