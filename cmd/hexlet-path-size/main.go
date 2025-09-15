package main

import (
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
				Aliases: []string{"h"},
				Usage:   "print sizes in human readable format (KB, MB, GB)",
				Value:   false,
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			recursiveFlag := cmd.Bool("recursive")
			fmt.Println("recursive:", recursiveFlag)
			allFlag := cmd.Bool("all")
			fmt.Println("all:", allFlag)
			humanFlag := cmd.Bool("human")
			fmt.Println("human:", humanFlag)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
