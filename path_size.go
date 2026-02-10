package code

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, all, human bool) (string, error) {
	var size int64
	var err error
	
	size, err = getSize(path, all, recursive)
	if err != nil {

		return "", err
	}

	if human {
		return humanReadable(size)
	}

	return fmt.Sprintf("%vB", size), nil
}

func getSize(path string, all, recursive bool) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if !fileInfo.IsDir() {
		return fileInfo.Size(), nil
	}

	size, err := readDirectory(path, all, recursive)
	if err != nil {
		return 0, fmt.Errorf("ошибка readDirectory в getSize: %v", err)
	}

	return size, nil
}

func readDirectory(path string, all, recursive bool) (int64, error) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return 1, errors.New("произошла ошибка чтения директории")
	}

	var allSize int64
	for _, entry := range dirEntries {
		isHidden := strings.HasPrefix(entry.Name(), ".")
		if !all && isHidden {
			continue
		}
		if entry.IsDir() {
			if !recursive {
				continue
			}
			//зайти внутрь и сумировать
			subDirSize, errReadDir := readDirectory(filepath.Join(path, entry.Name()), all, recursive)
			if errReadDir != nil {
				continue
			}
			allSize += subDirSize

		}

		if !entry.IsDir() {
			entryInfo, errEntryInfo := entry.Info()
			if errEntryInfo != nil {
				continue
			}
			allSize += entryInfo.Size()
		}

	}
	return allSize, nil
}

func humanReadable(size int64) (string, error) {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	i := 0
	sizeHuman := float64(size)
	for sizeHuman >= float64(1024) && i < len(units)-1 {
		sizeHuman = sizeHuman / float64(1024)
		i++
	}
	if i >= len(units) {

		return fmt.Sprintf("%.1f%s", sizeHuman, units[len(units)-1]), nil
	}
	return fmt.Sprintf("%.1f%s", sizeHuman, units[i]), nil
}
