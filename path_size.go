package pz

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func GetPathSize(path string, recursive, all, human bool) (string, error) {
	var size int64
	var err error

	if recursive {
		size, err = Recursive(path)
		if err != nil {
			return "", err
		}
		if human {
			return HumanReadable(size)
		}
		return strconv.FormatInt(size, 10), nil
	}

	if all {
		size, err = GetSizeAll(path)
		if err != nil {
			return "", err
		}
		if human {
			return HumanReadable(size)
		}
		return strconv.FormatInt(size, 10), nil
	}

	size, err = GetSize(path)
	if err != nil {
		return "", err
	}
	if human {
		return HumanReadable(size)
	}
	return strconv.FormatInt(size, 10), nil
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
		} else {
			totalSize += info.Size()
		}
	}
	return totalSize, nil
}

func GetSizeAll(path string) (int64, error) {
	var allSize int64

	entries, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if entries.IsDir() { //если не директория, то просто возвращает размер файла
		dirEntries, err := os.ReadDir(path)
		if err != nil {
			return 0, err
		}
		for _, entry := range dirEntries {
			if entry.IsDir() {
				continue // пропускаем поддиректории
			}
			size, err := GetSize(filepath.Join(path, entry.Name()))
			if err != nil {
				return 0, err
			}
			allSize += size
		}
	} else {
		// если это файл, возвращаем его размер
		return entries.Size(), nil
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
		if entry.IsDir() {
			subDirSize := ReadDirectory(filepath.Join(path, entry.Name()))
			allSize += subDirSize
		} else {
			size, err := GetSize(filepath.Join(path, entry.Name()))
			if err != nil {
				fmt.Println("ошибки появилась", err)
				continue
			}
			allSize += size
		}
	}

	return allSize
}

func HumanReadable(size int64) (string, error) {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	i := 0
	sizeHuman := float64(size)
	for sizeHuman >= float64(1024) && i < len(units)-1 {
		sizeHuman = sizeHuman / float64(1024)
		i++
	}
	if i >= len(units) {
		return fmt.Sprintf("%.2f%s", sizeHuman, units[len(units)-1]), nil
	}
	return fmt.Sprintf("%.2f%s", sizeHuman, units[i]), nil
}
