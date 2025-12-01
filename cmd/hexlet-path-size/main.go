package main

import (
	pz "code"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

func main() {
	var patc string
	fmt.Printf("Введите путь к файлу или директории: ")
	fmt.Scan(&patc)

	cmd := &cli.Command{
		Name:                   "hexlet-path-size",
		Usage:                  "print size of a file or directory",
		UseShortOptionHandling: true,

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "Calculates the size of ALL files in the directory (full recursion).",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "show size for all file in directory",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"h"},
				Usage:   "human-readable sizes (auto-select unit)",
				Value:   false,
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			res, err := pz.GetPathSize(filepath.Join(patc), cmd.Bool("recursive"), cmd.Bool("all"), cmd.Bool("human"))
			if err != nil {
				return err
			}
			fmt.Println(res)
			return err
		},
	}
	flag.Parsed()
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
