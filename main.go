package main

import (
	"log"
	"net/http"
)

func main() {
	InitDB()

	http.HandleFunc("/get-brands", GetBrandsHandler)
	http.HandleFunc("/get-models", GetModelsHandler)

	// Отдаём дашборд
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dashboard.html")
	})

	log.Println("Сервер запущен на 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
