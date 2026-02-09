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

	//if recursive {
	//	size, err = recursiveFile(path)
	//	if err != nil {
	//		return "", err
	//	}
	//	if human {
	//		return humanReadable(size)
	//	}
	//	return strconv.FormatInt(size, 10), nil
	//}
	//
	//if all {
	//	size, err = getSizeAll(path)
	//	if err != nil {
	//		return "", err
	//	}
	//	if human {
	//		return humanReadable(size)
	//	}
	//	return strconv.FormatInt(size, 10), nil
	//}

	size, err = getSize(path, all, recursive)
	if err != nil {

		return "", err
	}

	if human {
		return humanReadable(size)
	}

	//result :=
	return fmt.Sprintf("%vB", size), nil
}

//func recursiveFile(path string) (int64, error) {
//	entries, err := os.ReadDir(path)
//	if err != nil {
//		return 0, fmt.Errorf("ошибка чтения директории %s: %w", path, err)
//	}
//	var totalSize int64
//	for _, entry := range entries {
//		info, err := entry.Info()
//		if err != nil {
//			return 0, fmt.Errorf("ошибка чтения директории %s: %w", entry.Name(), err)
//		}
//		if entry.IsDir() {
//			subDirPatc := filepath.Join(path, entry.Name())
//			subDirSize, err := recursiveFile(subDirPatc)
//			if err != nil {
//				return 0, fmt.Errorf("ошибка чтения директории %s: %w", entry.Name(), err)
//			}
//			totalSize += subDirSize
//		} else {
//			totalSize += info.Size()
//		}
//	}
//	return totalSize, nil
//}

//func getSizeAll(path string) (int64, error) {
//	var allSize int64
//
//	fileInfo, err := os.Lstat(path)
//	if err != nil {
//		return 0, err
//	}
//
//	if !fileInfo.IsDir() {
//		return fileInfo.Size(), nil
//	}
//	//если не директория, то просто возвращает размер файла
//	dirEntries, err := os.ReadDir(path)
//	if err != nil {
//		return 0, err
//	}
//	for _, entry := range dirEntries {
//		if entry.IsDir() {
//			continue // пропускаем поддиректории
//		}
//		size, errgetSize := getSize(filepath.Join(path, entry.Name()))
//		if errgetSize != nil {
//			return 0, errgetSize
//		}
//		allSize += size
//	}
//
//	return allSize, nil
//}

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
		if all && isHidden {
			allDirectory, errAllDir := readDirectory(filepath.Join(path, entry.Name()), all, recursive)
			if errAllDir != nil {
				continue
			}
			allSize += allDirectory
			//continue // допустим у меня и то и то условие выполняется тогда нужно начинать читать скрытые директории и файлы
			// получается мне тут не continue нужен, а функция которая пойдет по скрытым файлам
			//getSize(path + "/" + entry.Name())
		}

		if entry.IsDir() {
			subDirSize, errReadDir := readDirectory(filepath.Join(path, entry.Name()), all, recursive)
			if errReadDir != nil {
				continue
			}
			allSize += subDirSize
		} else {
			size, errGetSize := getSize(filepath.Join(path, entry.Name()), all, recursive)
			if errGetSize != nil {
				continue
			}
			allSize += size
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
