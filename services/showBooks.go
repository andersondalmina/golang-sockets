package services

import (
	"fmt"

	"github.com/andersondalmina/golang-sockets/persist"
)

func DisplayBooks(books []persist.Book) {
	for _, item := range books {
		fmt.Printf("%d - %s\n", item.ID, item.Title)
	}
}
