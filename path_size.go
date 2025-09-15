package main

import (
	"fmt"
	"os"
)

func main() {
	var patc string
	fmt.Printf("Введите путь к файлу или директории: ")
	fmt.Scan(&patc)
	size, err := GetSize(patc)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(size, patc)
}

func GetSize(name string) (int64, error) {
	info, err := os.Stat(name)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
