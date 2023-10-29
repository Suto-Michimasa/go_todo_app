package controllers

import (
	"go_todo_app/app/models"
	"log"
	"net/http"
)

// Reactアプリケーションのビルドされたindex.htmlへのパスを定数として定義
const ReactAppIndexPath = "../../frontend/build/index.html"

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.ServeFile(w, r, "../../frontend/build/index.html")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		http.ServeFile(w, r, "../../frontend/build/index.html")
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if _, err := session(w, r); err == nil {
			// セッションが存在する場合、/todosにリダイレクト
			http.Redirect(w, r, "/todos", http.StatusSeeOther)
			return
		}
		http.ServeFile(w, r, ReactAppIndexPath)

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

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if _, err := session(w, r); err == nil {
			// セッションが存在する場合、/todosにリダイレクト
			http.Redirect(w, r, "/todos", http.StatusSeeOther)
			return
		}
		http.ServeFile(w, r, ReactAppIndexPath)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}

	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
