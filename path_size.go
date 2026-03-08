package code

import (
	"fmt"
	"os"
)

func GetPathSize(filePath string, human bool) (string, error) {
	var result string

	sizeBytes, err := GetSize(filePath)
	if err != nil {
		return "ошибка при чтении файла или директории", err
	}

	if human {
		result = fmt.Sprintf("%s\t%s", FormatSize(float64(sizeBytes)), filePath)
	} else {
		result = fmt.Sprintf("%d\t%s", sizeBytes, filePath)
	}

	return result, nil
}

func GetSize(filePath string) (int64, error) {
	var bytes int64 = 0

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

	return fmt.Sprintf("%.2f %s", bytes, units[i])
}
