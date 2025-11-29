# Инструкция для создания программы на Go о парфюме Tom Ford Lost Cherry

1. **Создайте новый проект**:
   - Создайте папку, например `tom_ford_perfume` и перейдите в неё в терминале.

2. **Сохраните код**:
   - Создайте файл `main.go` с приведённым ниже кодом:

    ```go
    package main

    import (
        "encoding/json"
        "fmt"
        "net/http"
    )

    type Perfume struct {
        Name        string   `json:"name"`
        Brand       string   `json:"brand"`
        Price       float64  `json:"price"`
        Notes       []string `json:"notes"`
        Availability bool     `json:"availability"`
    }

    func getPerfumeInfo() Perfume {
        return Perfume{
            Name:        "Lost Cherry",
            Brand:       "Tom Ford",
            Price:       250.00,
            Notes:       []string{"Cherry", "Almond", "Tobacco", "Vanilla"},
            Availability: true,
        }
    }

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
    ```

3. **Запустите сервер**:
   - В терминале выполните команду:
     ```bash
     go run main.go
     ```
   - Сервер будет доступен по адресу [http://localhost:8080/perfume](http://localhost:8080/perfume).

4. **Загрузите на GitHub**:
   - Инициализируйте git:
     ```bash
     git init
     ```
   - Добавьте файлы:
     ```bash
     git add main.go
     ```
   - Сделайте коммит:
     ```bash
     git commit -m "Initial commit for Tom Ford Lost Cherry perfume application"
     ```
   - Создайте репозиторий на GitHub и свяжите его с локальным:
     ```bash
     git remote add origin <URL вашего репозитория>
     ```
   - Отправьте код на GitHub:
     ```bash
     git push -u origin master
     ```

Теперь у вас есть программа о парфюме Tom Ford Lost Cherry, и она загружена на GitHub! Если возникнут вопросы, дайте знать!
