package code

import (
	"fmt"
	"os"
	"strings"
)

func GetPathSize(filePath string, human, all bool) (string, error) {
	var result string

	sizeBytes, err := GetSize(filePath, all)
	if err != nil {
		return "ошибка при чтении файла или директории", err
	}

	if human {
		result = fmt.Sprintf("%s\t%s", FormatSize(float64(sizeBytes)), filePath)
	} else {
		result = fmt.Sprintf("%dB\t%s", sizeBytes, filePath)
	}

	return result, nil
}

func GetSize(filePath string, all bool) (int64, error) {
	var bytes int64 = 0

	if !all && isHiddenFile(filePath) {
		return 0, fmt.Errorf("ошибка при чтении файла или директории: %w", os.ErrPermission)
	}

	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		return 0, fmt.Errorf("ошибка при чтении файла или директории: %w", err)
	}

	if fileInfo.IsDir() {
		files, err := os.ReadDir(filePath)
		if err != nil {
			return 0, fmt.Errorf("ошибка при чтении директории: %w", err)
		}

		for _, file := range files {
			if !file.IsDir() {
				if !all && isHiddenFile(file.Name()) {
					continue
				}
				info, err := file.Info()
				if err != nil {
					return 0, err
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
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	i := 0

	for bytes >= 1024 && i < len(units)-1 {
		bytes /= 1024
		i++
	}

	return fmt.Sprintf("%.2f%s", bytes, units[i])
}

func isHiddenFile(filePath string) bool {
	return strings.HasPrefix(filePath, ".")
}
