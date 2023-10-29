package controllers

import (
	"go_todo_app/config"
	"net/http"
	"path/filepath"
)

func StartMainServer() error {
	// Reactのビルドファイルが格納されているディレクトリへのパス
	staticPath := filepath.Join("frontend", "build")

	// /static/ でアクセスされたリクエストは、Reactの静的ファイルを提供
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(staticPath, "static")))))
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
