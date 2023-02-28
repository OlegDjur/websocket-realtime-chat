package db

import "database/sql"

type Database struct {
	db *sql.DB
}

func InitDB() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5432/go-chat?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
