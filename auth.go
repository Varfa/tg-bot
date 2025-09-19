package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Загружаем .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env, используем системные переменные")
	}
}

// Проверка логина и установка cookie
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	adminUser := os.Getenv("ADMIN_USERNAME")
	adminPass := os.Getenv("ADMIN_PASSWORD")

	if username == adminUser && password == adminPass {
		cookie := http.Cookie{
			Name:  "logged_in",
			Value: "true",
			Path:  "/",
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/dashboard.html", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// Проверка cookie при доступе к dashboard
func Authenticated(r *http.Request) bool {
	cookie, err := r.Cookie("logged_in")
	return err == nil && cookie.Value == "true"
}

// Logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "logged_in",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
