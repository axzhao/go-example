package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Article struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	CreatedDate sql.NullTime   `json:"create_date"`
	Body        sql.NullString `json:"body"`
	User        sql.NullInt64  `json:"user"`
}

func nullString(v string) sql.NullString {
	if v == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{Valid: true, String: v}
}

func ExampleSelect() {
	db, err := sql.Open("postgres", "postgres://postgres:password@127.0.0.1/postgres?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query(`SELECT * FROM article where id = $1`, 1)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var a Article
		if err := rows.Scan(&a.ID, &a.Title, &a.CreatedDate, &a.Body, &a.User); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("data: %#v\n", a)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
	db.Close()

	// Output:
	// 1
}
