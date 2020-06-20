package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ConnectionHandler holds database connection
type ConnectionHandler struct {
	connection *gorm.DB
}

// New creates a new Connection handler and opens a gorm connection
func New(username string, password string, host string, port string, database string) (*ConnectionHandler, error) {
	connection, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?parseTime=true",
			username,
			password,
			host,
			port,
			database,
		),
	)
	if err != nil {
		return &ConnectionHandler{}, err
	}

	return &ConnectionHandler{connection: connection}, nil
}

// GetConnection returns the connection handlers
func (connectionHandler ConnectionHandler) GetConnection() *gorm.DB {
	return connectionHandler.connection
}

// Close closes the db connection
func (connectionHandler *ConnectionHandler) Close() error {
	err := connectionHandler.connection.Close()

	return err
}
