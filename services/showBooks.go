package services

import (
	"fmt"

	"github.com/andersondalmina/golang-sockets/persist"
)

// DisplayBooks list all books returned from server
func DisplayBooks(books []persist.Book) {
	if len(books) == 0 {
		fmt.Println("No book found")
		return
	}

	for _, item := range books {
		fmt.Printf("%d - %s\n", item.ID, item.Title)
	}
}
