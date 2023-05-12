package clients

import "database/sql"

type DBClient interface {
	DB() *sql.DB
}
