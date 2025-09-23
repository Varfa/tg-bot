package main

import (
	"net/http"
	"os"
)

// Проверка авторизации
func Authenticated(r *http.Request) bool {
	cookie, err := r.Cookie("logged_in")
	return err == nil && cookie.Value == "true"
}

// Login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Отдаем файл index.html, а не login.html
		http.ServeFile(w, r, "/root/tg-bot/Index/static/index.html")
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
		// Редирект на корень — там отдаётся dashboard.html для авторизованных
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// При неправильном логине тоже отдаём страницу входа
	http.ServeFile(w, r, "/root/tg-bot/Index/static/index.html")
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
	// Редирект на страницу входа
	http.Redirect(w, r, "/static/index.html", http.StatusSeeOther)
}
