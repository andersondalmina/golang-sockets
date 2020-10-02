package cmd

import (
	"fmt"
	"os"

	"github.com/andersondalmina/golang-sockets/persist"
	"github.com/andersondalmina/golang-sockets/sockets"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server port",
	Short: "Initialize the TCP Server",
	Run: func(cmd *cobra.Command, args []string) {
		persist.CreateDatabase()

		fmt.Println("Socket server starting")

		server := sockets.CreateSocketServer(os.Getenv("SOCKET_HOST"), os.Getenv("SOCKET_PORT"))
		server.Listen()
	},
}
