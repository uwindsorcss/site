package models

type Markdown_page struct {
	Id    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Text  string `db:"text" json:"text"`
}
