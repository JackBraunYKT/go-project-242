package main

import (
	"code"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	isHumanFormat := false
	isAll := false
	isRecursive := false

	cmd := &cli.Command{
		Name: "hexlet-path-size",

		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "recursive",
				Aliases:     []string{"r"},
				Usage:       "recursive size of directories",
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:        "human",
				Aliases:     []string{"H"},
				Usage:       "human-readable sizes (auto-select unit)",
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:        "all",
				Aliases:     []string{"a"},
				Usage:       "include hidden files and directories",
				DefaultText: "false",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.NArg() == 0 {
				return fmt.Errorf("ошибка: не указан путь к файлу или директории")
			}

			path := cmd.Args().First()

			if cmd.Bool("human") {
				isHumanFormat = true
			}

			if cmd.Bool("all") {
				isAll = true
			}

			if cmd.Bool("recursive") {
				isRecursive = true
			}

			result, err := code.GetPathSize(path, isRecursive, isHumanFormat, isAll)
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s", result, path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "ошибка: %v\n", err)
		os.Exit(1)
	}
}
