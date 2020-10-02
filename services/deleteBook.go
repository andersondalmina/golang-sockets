package services

import (
	"github.com/andersondalmina/golang-sockets/persist"
	"github.com/manifoldco/promptui"

	"github.com/go-pg/pg/v10"
)

// DeleteBookMenu open delete book menu
func DeleteBookMenu() (action string, params string, err error) {
	prompt := promptui.Prompt{
		Label: "Book Code",
	}

	title, err := prompt.Run()
	checkError(err)

	return "deleteBookByTitle", title, nil
}

// DeleteBooksByTitle delete a book by the title given
func DeleteBooksByTitle(title string) error {
	book, err := persist.GetBookByTitle(title)
	if err == pg.ErrNoRows {
		return err
	} else if err != nil {
		panic(err)
	}

	return persist.DeleteBookByTitle(book)
}
