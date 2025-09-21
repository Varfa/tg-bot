package main

import (
	"log"
	"net/http"
)

func main() {
	InitDB()

	// Приватные API — проверка авторизации внутри handler
	http.HandleFunc("/get-brands", GetBrandsHandler)
	http.HandleFunc("/get-models", GetModelsHandler)

	// Логин и логаут
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)

	// Дашборд — только для авторизованных
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !Authenticated(r) {
			http.Redirect(w, r, "/login.html", http.StatusSeeOther)
			return
		}
		http.ServeFile(w, r, "dashboard.html")
	})

	log.Println("Сервер запущен на 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
