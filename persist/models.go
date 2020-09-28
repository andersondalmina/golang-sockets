package persist

// Book is the model for books
type Book struct {
	tableName struct{} `pg:"livros,alias:l"`

	ID    int64  `pg:"codigo"`
	Title string `pg:"titulo"`
}

// Author is the model for author
type Author struct {
	tableName struct{} `pg:"autor,alias:a"`

	ID   int64  `pg:"codigo"`
	Name string `pg:"nome"`
}
