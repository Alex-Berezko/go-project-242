package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"slices"
	"syscall"
)

func main() {
	var patc string
	fmt.
		fmt.Scan(&name)
	res := fmt.Sprintf(GetSize(name))
	io.WriteString(os.Stdout, res)

}

func GetSize(name string) (uint64, error) {
	Lstat(name, error())
	//if s == err() {
	//	s := Lstat(name)
	//
	//}
	return s, nil
}

func ReadDir(name string) ([]DirEntry, error) {
	// ReadDir читает именованный каталог, возвращая все его записи каталога,
	//отсортированные по имени файла.
	//Если происходит ошибка при чтении каталога,
	//ReadDir возвращает записи, которые он смог прочитать до ошибки вместе с ошибкой.
	f, err := openDir(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dirs, err := f.ReadDir(-1)
	slices.SortFunc(dirs, func(a, b DirEntry) int {
		return bytealg.CompareString(a.Name(), b.Name())
	})
	return dirs, err
}

func Lstat(name string) (FileInfo, error) {
	// Lstat возвращает FileInfo с описанием именованного файла. Если файл является символьной ссылкой,
	//возвращенный FileInfo описывает символьную ссылку. Lstat не пытается следовать по ссылке.
	//Если ошибка, то она будет типа *PathError.
	testlog.Stat(name)
	return lstatNolog(name)

}

//func GetPathSize(path string, recursive, human, all bool) (string, error) {
