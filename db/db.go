package db

import (
	"database/sql"
	"fmt"
	"net/url"
)

// Client is a database client
type Client struct {
	db *sql.DB
}

// New instantiates DB
func New(host string, port int32, username, password, dbname string) (*Client, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		username,
		url.QueryEscape(password),
		host,
		port,
		dbname,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Client{db: db}, nil
}

// Close closes the connection to the DB
func (c *Client) Close() error {
	return c.db.Close()
}
