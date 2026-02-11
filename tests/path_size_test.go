package tests

import (
	"code"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	testDir := t.TempDir()
	testFile := filepath.Join(testDir, "testFile.txt")
	testContent := []byte("Hello World, Hello Hexlet!")
	err := os.WriteFile(testFile, testContent, 0644)
	if err != nil {
		t.Fatalf("не удалось записать текст в файл %v", err)
	}
	size, errSize := code.GetPathSize(testFile, true, false, false)
	require.NoError(t, errSize)

	info, err := os.Stat(testFile)
	require.NoError(t, err)

	expectedSize := fmt.Sprintf("%dB", info.Size())

	require.Equal(t, expectedSize, size, "размер файла должен соответствовать реальному размеру")
}
