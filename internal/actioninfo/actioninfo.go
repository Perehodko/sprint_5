package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Парсим данные
		err := dp.Parse(data)
		if err != nil {
			// Логируем ошибку парсинга и переходим к следующей итерации
			log.Println("parse error:", err)
			continue
		}

		// Получаем информацию о активности
		info, err := dp.ActionInfo()
		if err != nil {
			// Логируем ошибку получения информации
			log.Println("action info error:", err)
			continue
		}

		// Выводим информацию
		fmt.Println(info)
	}
}
