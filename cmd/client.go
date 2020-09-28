package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use:   "client host:port",
	Short: "Initialize the TCP Client",
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) != 3 {
			fmt.Fprintf(os.Stderr, "Usage: %s %s host:port\n", os.Args[0], os.Args[1])
			os.Exit(1)
		}

		service := os.Args[2]

		tcpAddr, err := net.ResolveTCPAddr("tcp", service)
		checkError(err)

		fmt.Println(tcpAddr)

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		checkError(err)
		defer conn.Close()

		fmt.Println(conn)

		_, err = conn.Write([]byte("test\n"))
		checkError(err)

		reply := make([]byte, 1024)
		_, err = conn.Read(reply)
		checkError(err)

		fmt.Println(string(reply))
	},
}
