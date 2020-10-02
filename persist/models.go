package persist

// Book is the model for books
type Book struct {
	tableName struct{} `pg:"livros,alias:l"`

	ID       int64      `pg:"codigo,pk"`
	Title    string     `pg:"titulo"`
	Editions []*Edition `pg:"rel:has-many,join_fk:codigolivro"`
}

// Author is the model for author
type Author struct {
	tableName struct{} `pg:"autor,alias:a"`

	ID   int64  `pg:"codigo,pk"`
	Name string `pg:"nome"`
}

// Edition is the model for books editions
type Edition struct {
	tableName struct{} `pg:"edicao,alias:e"`

	Book *Book `pg:"rel:belongs-to,join_fk:codigo"`

	ID     int64  `pg:"codigolivro,pk"`
	Number int64  `pg:"numero"`
	Year   string `pg:"ano"`
}
