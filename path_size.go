package pz

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func GetPathSize(path string, recursive, all, human bool) (string, error) {
	var size int64
	if recursive || all {
		if recursive {
			
			size, err := Recursive(path)
			return strconv.Quote(strconv.FormatInt(size, 10)), err
		}
		if all {
			size, err := GetSizeAll(path)
			return strconv.Quote(strconv.FormatInt(size, 10)), err
		}
	}
	if human {
		res, err := HumanReadable(size)
		return res, err
	}
	size, err := GetSize(path)
	return strconv.Quote(strconv.FormatInt(size, 10)), err
}

func Recursive(path string) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("ошибка чтения директории %s: %w", path, err)
	}
	var totalSize int64
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return 0, fmt.Errorf("ошибка чтения директории %s: %w", entry.Name(), err)
		}
		if entry.IsDir() {
			subDirPatc := filepath.Join(path, entry.Name())
			subDirSize, err := Recursive(subDirPatc)
			if err != nil {
				return 0, fmt.Errorf("ошибка чтения директории %s: %w", entry.Name(), err)
			}
			totalSize += subDirSize
			return totalSize, nil
		}
		totalSize += info.Size()
	}
	return totalSize, nil
}

func GetSizeAll(path string) (int64, error) {
	var allSize, size int64

	entries, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if entries.IsDir() { //если не директория, то просто возвращает размер файла
		entries, err := os.ReadDir(path)
		if err != nil {
			return 0, err
		}
		for _, entry := range entries {
			if entry.IsDir() {
				return 0, err
			}
			size, err = GetSize(path + "/" + entry.Name())
		}
		allSize += size
	}
	return allSize, nil
}

func GetSize(path string) (int64, error) {

	entries, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if !entries.IsDir() {
		return entries.Size(), nil
	}
	size := ReadDirectory(path)
	return size, nil
}

func ReadDirectory(path string) int64 {

	entries, err := os.ReadDir(path)

	if err != nil {
		return 0
	}
	var allSize int64
	for _, entry := range entries {
		size, err := GetSize(path + "/" + entry.Name())
		allSize += size
		if err != nil {
			fmt.Println("ошибки появилась", err)
		}
		if entry.IsDir() {
			size := ReadDirectory(path + "/" + entry.Name())
			return size
		}
	}

	return allSize
}

func HumanReadable(size int64) (string, error) {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	i := 0
	sizeHuman := float64(size)
	for sizeHuman > float64(1024) {
		sizeHuman = sizeHuman / float64(1024)
		i++
	}
	if i > len(units) {
		return fmt.Sprintf("%.2f%s", sizeHuman, units[len(units)-1]), nil
	}
	return fmt.Sprintf("%.2f%s", sizeHuman, units[i]), nil

}
