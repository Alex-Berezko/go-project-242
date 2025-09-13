// path_size.go
package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	fmt.Print("Введите путь к файлу: ")
	fmt.Scan(&path)

	size, err := GetPathSize(path)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Размер файла: %d байт\n", size)
}

// GetPathSize возвращает размер файла в байтах
func GetPathSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
