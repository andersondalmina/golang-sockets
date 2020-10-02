package persist

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func CreateDatabase() {
	db = pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		fmt.Println("Unable to connecto to database")
		os.Exit(1)
	}
}

func GetBookByTitle(title string) (Book, error) {
	var book Book
	err := db.Model(&book).
		Where("titulo LIKE ?", title).
		Relation("Editions").
		Limit(1).
		Select()

	return book, err
}

func SearchBookByTitle(term string) []Book {
	var books []Book
	err := db.Model(&books).
		Where("titulo LIKE ?", "%"+term+"%").
		Order("titulo ASC").
		Select()

	if err != nil {
		panic(err)
	}

	return books
}

func SearchBookByYear(term string) {

}

func DeleteBookByTitle(book Book) error {
	editions := book.Editions

	_, err := db.Model(&editions).WherePK().Delete()
	if err != nil {
		panic(err)
	}

	_, err = db.Model(&book).WherePK().Delete()
	if err != nil {
		panic(err)
	}

	return nil
}

func UpdateBook() {

}
