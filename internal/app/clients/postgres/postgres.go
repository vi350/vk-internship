package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/vi350/vk-internship/internal/app/clients"
)

type Client struct {
	DBp *sql.DB
}

func New() clients.DBClient {
	DBp, err := connect()
	if err != nil {
		panic(err)
	}

	return &Client{
		DBp: DBp,
	}
}

func (c *Client) DB() *sql.DB {
	return c.DBp
}
