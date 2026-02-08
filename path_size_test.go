package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"code"
)

func TestGetPathSize_File(t *testing.T) {
	// Тест для одного файла - проверяем, что размер соответствует реальному размеру файла
	testFile := filepath.Join("..", "testdata", "this_file.md")

	// Получаем размер файла через GetSize()
	size, err := code.GetSize(testFile)
	require.NoError(t, err)

	// Получаем реальный размер файла для проверки
	info, err := os.Stat(testFile)
	require.NoError(t, err)
	expectedSize := info.Size()

	// Убеждаемся, что результат соответствует сумме размеров (для файла это просто размер файла)
	require.Equal(t, expectedSize, size, "размер файла должен соответствовать реальному размеру")
}
