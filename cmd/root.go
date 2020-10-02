package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/andersondalmina/golang-sockets/persist"
	"github.com/andersondalmina/golang-sockets/services"
	"github.com/andersondalmina/golang-sockets/sockets"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "TCP Socket university project",
	Run: func(cmd *cobra.Command, args []string) {
		action, params, err := openMenu()
		checkError(err)

		socketClient := sockets.CreateSocketClient(os.Getenv("SOCKET_HOST"), os.Getenv("SOCKET_PORT"))
		socketClient.Send(action, params)

		handleSocketData(action, socketClient.Reply)

		socketClient.Close()
	},
}

// Execute the command line app
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func openMenu() (action string, params string, err error) {
	prompt := promptui.Select{
		Label: "Select your action",
		Items: []string{"Create a book", "Update a Book", "Search a book", "Delete a book", "Exit"},
	}

	_, result, err := prompt.Run()
	checkError(err)

	return handleMenu(result)
}

func handleMenu(item string) (action string, params string, err error) {
	switch item {
	case "Create a book":
		fmt.Printf("Book created\n")
	case "Update a Book":
		fmt.Printf("Book Updated\n")
	case "Search a book":
		return services.SearchBookMenu()
	case "Delete a book":
		return services.DeleteBookMenu()
	case "Exit":
		os.Exit(0)
	}

	return "", "", nil
}

func handleSocketData(action string, r sockets.SocketReply) {
	switch action {
	case "searchBookByTitle":
	case "searchBookByAuthor":
	case "searchBookByYear":
	case "searchBookByNumber":
		var books []persist.Book
		err := json.Unmarshal(r.Data, &books)
		checkError(err)

		services.DisplayBooks(books)

	case "deleteBookByTitle":
		if r.Status == 404 {
			fmt.Println("Book not found")
			return
		}

		fmt.Println("Book removed successfully")
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
