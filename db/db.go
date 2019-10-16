package db

import (
    "fmt"

    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Connection struct {
    db *gorm.DB
}

func Open(username string, password string, host string, port string, database string) *Connection, error {
    return &Connection{db: gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", username, password, host, port, database))}
}

func (connection *Connection) Close() error {
    err := connection.db.Close()
    if err != nil {
	return err
    }

    return nil
}

func (connection *Connection) AutoMigration(values ...interface{}) *Connection {
    connection.db.AutoMigration(values)

    return connection
}
