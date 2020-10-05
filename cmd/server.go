package cmd

import (
	"fmt"

	"github.com/andersondalmina/golang-sockets/persist"
	"github.com/andersondalmina/golang-sockets/sockets"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server [host] [port]",
	Short: "Initialize the TCP Server",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		persist.CreateDatabase()

		fmt.Println("Socket server starting")

		server := sockets.CreateSocketServer(args[0], args[1])
		server.Listen()
	},
}
