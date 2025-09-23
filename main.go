package main

import (
	"log"
	"net/http"
	"os"
)

// Проверка авторизации
func Authenticated(r *http.Request) bool {
	cookie, err := r.Cookie("logged_in")
	return err == nil && cookie.Value == "true"
}

// LoginHandler — обработчик логина
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Отдаем страницу входа (index.html)
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
		http.Redirect(w, r, "/", http.StatusSeeOther) // Редирект на главную
		return
	}

	// При неправильном логине — снова страница входа
	http.ServeFile(w, r, "/root/tg-bot/Index/static/index.html")
}

// LogoutHandler — обработчик логаута
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "logged_in",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/static/index.html", http.StatusSeeOther)
}

func main() {
	// Отдача статики по /static/
	fs := http.FileServer(http.Dir("/root/tg-bot/Index/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Логин и логаут
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)

	// Главная страница / дашборд
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !Authenticated(r) {
			http.Redirect(w, r, "/static/index.html", http.StatusSeeOther) // Если неавторизован — редирект на вход
			return
		}
		// Авторизован — отдаем dashboard.html
		http.ServeFile(w, r, "/root/tg-bot/Index/static/dashboard.html")
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
