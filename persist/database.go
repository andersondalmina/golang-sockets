package persist

import (
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
}

func SearchBook(term string) []Book {
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

func DeleteBook() {

}

func UpdateBook() {

}
