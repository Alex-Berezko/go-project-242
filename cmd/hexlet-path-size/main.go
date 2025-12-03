package main

import (
	pz "code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
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
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
				Value:   false,
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args()
			if args.Len() == 0 {
				return fmt.Errorf("путь к файлу или директории не указан")
			}
			path := args.Get(0)

			res, err := pz.GetPathSize(path, cmd.Bool("recursive"), cmd.Bool("all"), cmd.Bool("human"))
			if err != nil {
				return err
			}
			fmt.Println(res)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
