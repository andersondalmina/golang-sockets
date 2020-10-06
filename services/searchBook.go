package services

import (
	"errors"
	"strconv"

	"github.com/andersondalmina/golang-sockets/persist"
	"github.com/manifoldco/promptui"
)

// SearchBookMenu open search book menu
func SearchBookMenu() (action string, params map[string]string, err error) {
	prompt := promptui.Select{
		Label: "Search a book",
		Items: []string{"By title", "By author", "By year", "By edition"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", nil, err
	}

	return handleSearchBookMenu(result)
}

func handleSearchBookMenu(item string) (action string, params map[string]string, err error) {
	switch item {
	case "By title":
		prompt := promptui.Prompt{
			Label: "Title",
		}

		title, err := prompt.Run()
		params := map[string]string{
			"title": title,
		}

		return "searchBookByTitle", params, err
	case "By author":
		prompt := promptui.Prompt{
			Label: "Author",
		}

		author, err := prompt.Run()
		params := map[string]string{
			"author": author,
		}

		return "searchBookByAuthor", params, err
	case "By year":
		prompt := promptui.Prompt{
			Label: "Published Year",
			Validate: func(input string) error {
				_, err := strconv.ParseFloat(input, 64)
				if err != nil {
					return errors.New("Invalid number")
				}
				return nil
			},
		}

		year, err := prompt.Run()
		params := map[string]string{
			"year": year,
		}

		return "searchBookByYear", params, err
	case "By edition":
		prompt := promptui.Prompt{
			Label: "Book Edition",
			Validate: func(input string) error {
				_, err := strconv.ParseFloat(input, 64)
				if err != nil {
					return errors.New("Invalid number")
				}
				return nil
			},
		}

		edition, err := prompt.Run()
		params := map[string]string{
			"edition": edition,
		}

		return "searchBookByEdition", params, err
	}

	return "", nil, nil
}

func SearchBooksByTitle(title string) []persist.Book {
	return persist.SearchBookByTitle(title)
}

func SearchBooksByAuthor(author string) []persist.Book {
	return persist.SearchBookByAuthor(author)
}

func SearchBooksByYear(year string) []persist.Book {
	return persist.SearchBookByYear(year)
}

func SearchBooksByEdition(edition string) []persist.Book {
	return persist.SearchBooksByEdition(edition)
}
