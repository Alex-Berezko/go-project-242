package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	// Тест для одного файла - проверяем, что размер соответствует реальному размеру файла
	testFile := filepath.Join("..", "path_size.go", "README.md")

	// Получаем размер файла через GetSize()
	size, err := GetPathSize(testFile, true, false, false)
	require.NoError(t, err)

	// Получаем реальный размер файла для проверки
	info, err := os.Stat(testFile)
	require.NoError(t, err)
	expectedSize := info.Size()

	// Убеждаемся, что результат соответствует сумме размеров (для файла это просто размер файла)
	require.Equal(t, expectedSize, size, "размер файла должен соответствовать реальному размеру")
}
