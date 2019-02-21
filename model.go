package api

import (
	"database/sql"
)

type DB struct {
	*sql.DB
}

type Resolver struct {
	DB *DB
}
