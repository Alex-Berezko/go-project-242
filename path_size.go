package pz

import (
	"fmt"
	"os"
	"strconv"
)

func GetPathSize(path string, recursive, all, human bool) (string, error) {
	var size int64
	if recursive || all {
		if recursive {
			size, err := GetSize(path)
			return strconv.Quote(strconv.FormatInt(size, 10)), err
		}
		if all {
			size, err := GetSize(path)
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

func GetSize(path string) (int64, error) {
	entries, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if !entries.IsDir() {
		//fmt.Printf("размер файла %s\n ", path)
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
		//fmt.Println("сценарий по проходу на файлам")
		size, err := GetSize(path + "/" + entry.Name())
		if err != nil {
			fmt.Println("ошибки появилась", err)
		}
		//fmt.Printf("size = %d, path = %s \n", size, path)
		if !entry.IsDir() {
			size, _ := GetSize(path + "/" + entry.Name())
			return size
		}
		allSize += size
	}
	fmt.Println("all= ", allSize)

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
