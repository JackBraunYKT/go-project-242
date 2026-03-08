package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(filePath string, recursive, human, all bool) (string, error) {
	var result string

	sizeBytes, err := GetSize(filePath, recursive, all)
	if err != nil {
		return "ошибка при чтении файла или директории", err
	}

	if human {
		result = FormatSize(float64(sizeBytes))
	} else {
		result = fmt.Sprintf("%dB", sizeBytes)
	}

	return result, nil
}

func GetSize(filePath string, recursive, all bool) (int64, error) {
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
					dirBytes, err := GetSize(fullPath, recursive, all)
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

func FormatSize(bytes float64) string {
	if bytes < 1024 {
		return fmt.Sprintf("%dB", int64(bytes))
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	i := 0

	for bytes >= 1024 && i < len(units)-1 {
		bytes /= 1024
		i++
	}

	return fmt.Sprintf("%.1f%s", bytes, units[i])
}

func isHiddenFile(filePath string) bool {
	return strings.HasPrefix(filePath, ".")
}
