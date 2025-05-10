package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Хендлер для отправки HTML формы
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что запрос именно к корневому пути
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Устанавливаем правильный Content-Type для HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Открываем файл index.html из корневой директории проекта
	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "Не удалось найти HTML шаблон: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Читаем содержимое файла
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Не удалось прочитать содержимое HTML шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Устанавливаем статус 200 OK и отправляем содержимое файла
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

// Хендлер для загрузки и обработки файла
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса - POST
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Парсим файл из формы
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Не удалось распарсить форму", http.StatusInternalServerError)
		return
	}

	// Получаем файл из формы
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Не удалось получить файл из формы", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Читаем данные из файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Не удалось прочитать данные из файла", http.StatusInternalServerError)
		return
	}

	// Передаем данные в функцию автоопределения и конвертации
	result, err := service.ConvertData(string(data))
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаем безопасное имя файла с использованием Unix-времени
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	fileExt := filepath.Ext(header.Filename)
	if fileExt == "" {
		fileExt = ".txt" // добавляем расширение, если его нет
	}
	outputFileName := "result_" + timestamp + fileExt

	// Создаем папку results, если её нет
	if err := os.MkdirAll("results", 0755); err != nil {
		http.Error(w, "Не удалось создать папку для результатов", http.StatusInternalServerError)
		return
	}

	// Сохраняем в папку results
	outputPath := filepath.Join("results", outputFileName)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		http.Error(w, "Не удалось создать выходной файл: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Записываем результат в файл
	_, err = outputFile.WriteString(result)
	if err != nil {
		http.Error(w, "Не удалось записать данные в выходной файл", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type для текстового ответа
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Возвращаем результат конвертации
	fmt.Fprint(w, result)
}
