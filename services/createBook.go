package services

import (
	"errors"
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
		Validate: func(input string) error {
			if len(input) > 1 {
				return errors.New("Max length 1")
			}

			_, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return errors.New("Invalid number")
			}
			return nil
		},
	}

	edition, err := prompt.Run()

	prompt = promptui.Prompt{
		Label: "Book Year",
		Validate: func(input string) error {
			if len(input) > 4 {
				return errors.New("Max length 4")
			}

			_, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return errors.New("Invalid number")
			}
			return nil
		},
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
