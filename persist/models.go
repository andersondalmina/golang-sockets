package persist

// Book is the model for books
type Book struct {
	tableName struct{} `pg:"livrostemp,alias:l"`

	ID      int64  `pg:"codigo,pk"`
	Title   string `pg:"titulo"`
	Author  string `pg:"autor"`
	Edition int64  `pg:"edicao"`
	Year    int64  `pg:"ano"`
}
