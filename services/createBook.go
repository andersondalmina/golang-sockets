package services

import (
	"strconv"

	"github.com/andersondalmina/golang-sockets/persist"
	"github.com/manifoldco/promptui"
)

func CreateBookMenu() (string, map[string]string, error) {
	prompt := promptui.Prompt{
		Label: "Book Title",
	}

	title, err := prompt.Run()

	prompt = promptui.Prompt{
		Label: "Book Author",
	}

	author, err := prompt.Run()

	prompt = promptui.Prompt{
		Label: "Book Edition",
	}

	edition, err := prompt.Run()

	prompt = promptui.Prompt{
		Label: "Book Year",
	}

	year, err := prompt.Run()

	params := map[string]string{
		"title":   title,
		"author":  author,
		"edition": edition,
		"year":    year,
	}

	return "createBook", params, err
}

func CreateBook(p map[string]string) (persist.Book, error) {
	edition, err := strconv.Atoi(p["edition"])
	if err != nil {
		panic(err)
	}

	year, err := strconv.Atoi(p["year"])
	if err != nil {
		panic(err)
	}

	return persist.CreateBook(p["title"], p["author"], int64(edition), int64(year))
}
