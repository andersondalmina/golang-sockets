package cmd

import (
	"fmt"
	"os"

	"github.com/andersondalmina/golang-sockets/persist"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search books by term",
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) != 3 {
			fmt.Fprintf(os.Stderr, "Usage: %s %s term\n", os.Args[0], os.Args[1])
			os.Exit(1)
		}

		term := os.Args[2]

		books := persist.SearchBook(term)

		for i := range books {
			fmt.Println(books[i].Title)
		}
	},
}
