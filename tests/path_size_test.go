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

	size, err := code.GetPathSize(testFile, true, false, false)
	require.NoError(t, err)

	info, err := os.Stat(testFile)
	require.NoError(t, err)

	expectedSize := fmt.Sprintf("%dB", info.Size())

	require.Equal(t, expectedSize, size, "размер файла должен соответствовать реальному размеру")
}
