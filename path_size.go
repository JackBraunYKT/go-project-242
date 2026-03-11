package code

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize возвращает размер файла или директории.
// Если recursive равен true, размер директорий вычисляется рекурсивно.
// Если human равен true, результат возвращается в человекочитаемом формате.
// Если all равен false, скрытые файлы и директории не учитываются.
func GetPathSize(filePath string, recursive, human, all bool) (string, error) {
	var result string

	sizeBytes, err := getSize(filePath, recursive, all)
	if err != nil {
		return "ошибка при чтении файла или директории", err
	}

	if human {
		result = formatSize(float64(sizeBytes))
	} else {
		result = fmt.Sprintf("%dB", sizeBytes)
	}

	return result, nil
}

func getSize(filePath string, recursive, all bool) (int64, error) {
	var bytes int64 = 0

	if !all && isHiddenFile(filepath.Base(filePath)) {
		return 0, nil
	}

	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		return 0, nil
	}

	if fileInfo.IsDir() {
		entries, err := os.ReadDir(filePath)
		if err != nil {
			return 0, nil
		}

		for _, entry := range entries {
			fullPath := filepath.Join(filePath, entry.Name())

			if entry.IsDir() {
				if recursive {
					dirBytes, err := getSize(fullPath, recursive, all)
					if err != nil {
						continue
					}
					bytes += dirBytes
				}
			} else {
				if !all && isHiddenFile(entry.Name()) {
					continue
				}
				info, err := entry.Info()
				if err != nil {
					continue
				}
				bytes += info.Size()
			}
		}
	} else {
		bytes = fileInfo.Size()
	}

	return bytes, nil
}

var units = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

func formatSize(bytes float64) string {
	if bytes < 1024 {
		return fmt.Sprintf("%dB", int64(bytes))
	}
	i := int(math.Log2(bytes) / 10)

	if i >= len(units) {
		i = len(units) - 1
	}

	value := bytes / math.Pow(1024, float64(i))
	return fmt.Sprintf("%.1f%s", value, units[i])
}

func isHiddenFile(filePath string) bool {
	return strings.HasPrefix(filePath, ".")
}
