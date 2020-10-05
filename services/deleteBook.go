package services

import (
	"github.com/andersondalmina/golang-sockets/persist"
	"github.com/manifoldco/promptui"
)

// DeleteBookMenu open delete book menu
func DeleteBookMenu() (string, map[string]string, error) {
	prompt := promptui.Prompt{
		Label: "Book Title",
	}

	title, err := prompt.Run()
	params := map[string]string{
		"title": title,
	}

	return "deleteBookByTitle", params, err
}

// DeleteBooksByTitle delete a book by the title given
func DeleteBooksByTitle(title string) error {
	return persist.DeleteBooksByTitle(title)
}
