package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Получаем бренды (только для авторизованных)
func GetBrandsHandler(w http.ResponseWriter, r *http.Request) {
	if !Authenticated(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	rows, err := db.Query("SELECT id, name FROM car_brands ORDER BY name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var brands []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		brands = append(brands, map[string]interface{}{"id": id, "name": name})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brands)
}

// Получаем модели по бренду (только для авторизованных)
func GetModelsHandler(w http.ResponseWriter, r *http.Request) {
	if !Authenticated(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	brandIDStr := r.URL.Query().Get("brandId")
	brandID, err := strconv.Atoi(brandIDStr)
	if err != nil || brandID == 0 {
		http.Error(w, "Неверный бренд ID", http.StatusBadRequest)
		return
	}

	rows, err := db.Query("SELECT id, name FROM car_models WHERE brand_id=$1 ORDER BY name", brandID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var models []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		models = append(models, map[string]interface{}{"id": id, "name": name})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models)
}
