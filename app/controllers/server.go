package controllers

import (
	"fmt"
	"go_todo_app/app/models"
	"go_todo_app/config"
	"net/http"
	"path/filepath"
)

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

func StartMainServer() error {
	// Reactのビルドファイルが格納されているディレクトリへのパス
	staticPath := filepath.Join("frontend", "build")

	// /static/ でアクセスされたリクエストは、Reactの静的ファイルを提供
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(staticPath, "static")))))
	http.HandleFunc("top", top)
	http.HandleFunc("index", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)

	// それ以外のリクエストは、Reactのindex.htmlを提供
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticPath, "index.html"))
	})

	// サーバを起動
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
