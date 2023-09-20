package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "example.txt"

	// Установка прав доступа 666
	err := os.Chmod(fileName, 0666)
	if err != nil {
		fmt.Println("Ошибка при изменении прав доступа:", err)
		return
	}

	fmt.Println("Права доступа к файлу успешно изменены.")
}
