package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Perfume представляет собой структуру для хранения информации о парфюме
type Perfume struct {
	Name       string  `json:"name"`
	Brand      string  `json:"brand"`
	Price      float64 `json:"price"`
	Notes      []string `json:"notes"`
	Availability bool   `json:"availability"`
}

// getPerfumeInfo возвращает информацию о парфюме Tom Ford Lost Cherry
func getPerfumeInfo() Perfume {
	return Perfume{
		Name:       "Lost Cherry",
		Brand:      "Tom Ford",
		Price:      250.00, // цена в условных единицах
		Notes:      []string{"Cherry", "Almond", "Tobacco", "Vanilla"},
		Availability: true,
	}
}

// handlePerfumeInfo выводит информацию о парфюме в формате JSON
func handlePerfumeInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	perfume := getPerfumeInfo()
	json.NewEncoder(w).Encode(perfume)
}

func main() {
	http.HandleFunc("/perfume", handlePerfumeInfo)
	fmt.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
