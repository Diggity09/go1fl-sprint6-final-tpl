package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Функция для автоматического определения формата данных и конвертации
func ConvertData(data string) (string, error) {
	if data == "" {
		return "", errors.New("Пустые данные")
	}

	// Очищаем строку от начальных и конечных пробелов
	cleanData := strings.TrimSpace(data)

	// Используем ContainsFunc для проверки наличия символов, не относящихся к коду Морзе
	containsNonMorseChar := strings.ContainsFunc(cleanData, func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	})

	// Если строка не содержит символов, которые не являются частью кода Морзе,
	// и в ней есть хотя бы одна точка или тире, считаем её кодом Морзе
	isMorse := !containsNonMorseChar && (strings.Contains(cleanData, ".") || strings.Contains(cleanData, "-"))

	if isMorse {
		// Конвертируем из кода Морзе в текст
		result := morse.ToText(cleanData)
		return result, nil
	} else {
		// Конвертируем из текста в код Морзе
		result := morse.ToMorse(cleanData)
		return result, nil
	}
}
