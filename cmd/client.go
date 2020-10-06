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

var socketClient sockets.SocketClient

func init() {
	rootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use:   "client [host] [port]",
	Short: "Initialize the TCP Client",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		for {
			socketClient = sockets.CreateSocketClient(args[0], args[1])

			action, params, err := openMenu()
			checkError(err)

			socketClient.Send(action, params)

			handleSocketData(action, socketClient.Reply)
		}
	},
}

func openMenu() (action string, params map[string]string, err error) {
	prompt := promptui.Select{
		Label: "Select your action",
		Items: []string{"Create a book", "Update a Book", "Search a book", "Delete a book", "Exit"},
	}

	_, result, err := prompt.Run()
	checkError(err)

	return handleMenu(result)
}

func handleMenu(item string) (action string, params map[string]string, err error) {
	switch item {
	case "Create a book":
		return services.CreateBookMenu()
	case "Update a Book":
		return services.UpdateBookMenu()
	case "Search a book":
		return services.SearchBookMenu()
	case "Delete a book":
		return services.DeleteBookMenu()
	case "Exit":
		socketClient.Close()
		os.Exit(0)
	}

	return "", nil, nil
}

func handleSocketData(action string, r sockets.SocketReply) {
	socketClient.Close()

	if r.Error != nil {
		fmt.Printf("Error: %s\n\n", r.Error)
		return
	}

	switch action {
	case "createBook":
		fmt.Println("Book created successfully")

	case "updateBook":
		fmt.Println("Book updated successfully")

	case "searchBookByTitle", "searchBookByAuthor", "searchBookByYear", "searchBookByEdition":
		var books []persist.Book
		err := json.Unmarshal(r.Data, &books)
		checkError(err)

		services.DisplayBooks(books)

	case "deleteBookByTitle":
		fmt.Println("Book removed successfully")
	}

	fmt.Println()
}
