package cmd

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server port",
	Short: "Initialize the TCP Server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) != 3 {
			fmt.Fprintf(os.Stderr, "Usage: %s %s port\n", os.Args[0], os.Args[1])
			os.Exit(1)
		}

		port := os.Args[2]
		tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
		checkError(err)

		listener, err := net.ListenTCP("tcp", tcpAddr)
		checkError(err)

		fmt.Printf("Server listening at port %s\n", port)

		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}

			handleClient(conn)
		}
	},
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}
