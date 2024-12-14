package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var count1 int = 0

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // Устанавливаем Content-Type

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(count1)))
	case "POST":
		r.ParseForm()
		s := r.FormValue("count")
		if s == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
			return
		}
		number, err := strconv.Atoi(s)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
			return
		}
		count1 += number
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Метод не поддерживается"))
	}
}

func main() {
	http.HandleFunc("/count", handler) // Обработчик для пути /count

	fmt.Println("Сервер запущен на http://localhost:3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
