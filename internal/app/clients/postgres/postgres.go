package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Client struct {
	DB *sql.DB
}

func New() *Client {
	DB, err := connect()
	if err != nil {
		panic(err)
	}

	return &Client{
		DB: DB,
	}
}
