package code

import (
	"fmt"
	"os"
)

func GetSize(filePath string) (string, error) {
	var resultSize int64 = 0

	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		return "", fmt.Errorf("ошибка при чтении файла или директории: %w", err)
	}

	if fileInfo.IsDir() {
		files, err := os.ReadDir(filePath)
		if err != nil {
			return "", fmt.Errorf("ошибка при чтении директории: %w", err)
		}

		for _, file := range files {
			if !file.IsDir() {
				info, err := file.Info()
				if err != nil {
					return "", err
				}
				resultSize += info.Size()
			}
		}
	} else {
		resultSize = fileInfo.Size()
	}

	return fmt.Sprintf("%dB\t%s", resultSize, filePath), nil
}
