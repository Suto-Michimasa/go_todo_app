package controllers

import (
	"go_todo_app/app/models"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "path_to_signup.html") // signup.htmlへのパスを適切に設定してください

	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse the form.", http.StatusBadRequest)
			log.Println(err)
			return
		}

		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}

		if err := user.CreateUser(); err != nil {
			http.Error(w, "Failed to create user.", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound) // 302: StatusFound

	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}
