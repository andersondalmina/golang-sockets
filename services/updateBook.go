package services

import (
	"strconv"

	"github.com/andersondalmina/golang-sockets/persist"
	"github.com/manifoldco/promptui"
)

func UpdateBookMenu() (string, map[string]string, error) {
	prompt := promptui.Prompt{
		Label: "Book Code",
	}

	id, err := prompt.Run()

	prompt = promptui.Prompt{
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
		"id":      id,
		"title":   title,
		"author":  author,
		"edition": edition,
		"year":    year,
	}

	return "updateBook", params, err
}

func UpdateBook(p map[string]string) (persist.Book, error) {
	id, err := strconv.Atoi(p["id"])
	if err != nil {
		panic(err)
	}

	edition, err := strconv.Atoi(p["edition"])
	if err != nil {
		panic(err)
	}

	year, err := strconv.Atoi(p["year"])
	if err != nil {
		panic(err)
	}

	return persist.UpdateBook(int64(id), p["title"], p["author"], int64(edition), int64(year))
}
