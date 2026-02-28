package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.NArg() == 0 {
				return fmt.Errorf("Ошибка: не указан путь к файлу или директории")
			}

			path := c.Args().First()

			size, err := GetSize(path)
			if err != nil {
				return err
			}

			fmt.Println(size)
			return nil
		},
	}

	cmd.Run(context.Background(), os.Args)
}

func GetSize(filePath string) (string, error) {
	var resultSize int64 = 0

	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		return "", fmt.Errorf("Ошибка при чтении файла или директории: %w", err)
	}

	if fileInfo.IsDir() {
		files, err := os.ReadDir(filePath)
		if err != nil {
			return "", fmt.Errorf("Ошибка при чтении директории: %w", err)
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
