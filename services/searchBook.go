package services

import (
	"github.com/andersondalmina/golang-sockets/persist"
	"github.com/manifoldco/promptui"
)

// SearchBookMenu open search book menu
func SearchBookMenu() (action string, params string, err error) {
	prompt := promptui.Select{
		Label: "Search a book",
		Items: []string{"By title", "By author", "By year", "By code"},
	}

	_, result, err := prompt.Run()
	checkError(err)

	return handleSearchBookMenu(result)
}

func handleSearchBookMenu(item string) (action string, params string, err error) {
	switch item {
	case "By title":
		prompt := promptui.Prompt{
			Label: "Title",
		}

		title, err := prompt.Run()
		checkError(err)

		return "searchBookByTitle", title, nil
	case "By author":
		prompt := promptui.Prompt{
			Label: "Author",
		}

		author, err := prompt.Run()
		checkError(err)

		return "searchBookByAuthor", author, nil
	case "By year":
		prompt := promptui.Prompt{
			Label: "Published Year",
		}

		year, err := prompt.Run()
		checkError(err)

		return "searchBookByYear", year, nil
	case "By code":
		prompt := promptui.Prompt{
			Label: "Book Code",
		}

		number, err := prompt.Run()
		checkError(err)

		return "searchBookByCode", number, nil
	}

	return "", "", nil
}

func SearchBooksByTitle(title string) []persist.Book {
	return persist.SearchBookByTitle(title)
}
