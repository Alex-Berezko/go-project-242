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
	//sizeHuman := FormatSize(GetSize(size))
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(size, patc)
}

func GetSize(path string) (int64, error) {

	info, err := os.Lstat(path)
	fmt.Println(info.IsDir())

	if info.IsDir() == false { //тут если это НЕ директория то возвращаем размер файла
		fmt.Printf("%s - это файл , его размер %v\n", path, uint((info.Size())))
		return info.Size(), nil

	}

	if info.IsDir() == true { // если это директория нужно это значение передать в функцию для расчета суммы файлов
		fmt.Printf("%s - это директория \n", path)

		size, err := ReadDir(path)
		fmt.Printf("это вывод этой команды size - %v, значение err -  %v \n", size, err)
	}
	if err != nil {
		fmt.Printf("Произошла ошибка \n")
		return 0, err
	}

	return info.Size(), nil
}

func ReadDir(name string) (int, error) { // по идеи эта функция должна посмотреть директорию
	entries, err := os.ReadDir(name)

	if err != nil {
		return 0, err
	}
	fmt.Printf("это колличество суммарное папок и файлов в данной дирректории - %v \n", len(entries))
	fmt.Printf("дальше файлы с отсеянными папками, но остались скрытые файлы \n")
	for _, file := range entries {
		if file.IsDir() != true { //отсеяли директории, и видны все файлы, в том числе и скрытые
			info, err := file.Info()
			if err != nil {
				return 0, err
				continue
			}
			//size += info.Size()
			//fmt.Println(size)
			fmt.Printf("Файл - %s, его размер - %v \n", file.Name(), info.Size())
			//totalSize += size
			//fmt.Printf("общий размер папки - %d", )

		}
	func FormatSize(size, *human) int {
		if size < 0 {27}
		}
		// а тут должны быть имена всех файлов
		// допустим на этом этапе вывели содержимое директории, но тут есть папка которая мне не нужна
		//fmt.Println(file.Info())
	}
	return len(entries), err

	c := fmt.Sprintf()
	//func Recursive() { //отсеять файлы с точкой в начале (скрытые файлы)
	//	entries, err := os.ReadDir(".")
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	for _, entry := range entries {
	//		// Пропускаем файлы, которые начинаются с точки
	//		if strings.HasPrefix(entry.Name(), ".") {
	//			continue
	//		}
	//
	//		fmt.Println(entry.Name())
	//	}
	//}

	//func FormatSize(size ) (int64) {
	//
	//}

	//if err != nil {
	//	return nil, err
	//}
	//for _, fi := range f {
	//	if fi.IsDir != true {
	//
	//	}
	//}
	//defer f.Close()
	//
	//dirs, err := f.ReadDir(-1)
	//slices.SortFunc(dirs, func(a, b DirEntry) int {
	//	return bytealg.CompareString(a.Name(), b.Name())
	//})
	//return dirs, err
}
