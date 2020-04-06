package db

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type ConnectionFactory struct {
    username string
    password string
    host string
    port string
    database string
    connection *gorm.DB
}

func NewConnectionFactory(username string, password string, host string, port string, database string) *ConnectionFactory {
    return &ConnectionFactory{
	username: username,
	password: password,
	host: host,
	port: port,
	database: database,
    }
}

func (connectionFactory *ConnectionFactory) GetConnection() *gorm.DB {
    if connectionFactory.connection != nil {
	return connectionFactory.connection
    }

    connection, err := gorm.Open(
	"mysql",
	fmt.Sprintf(
	    "%s:%s@(%s:%s)/%s?parseTime=true",
	    connectionFactory.username,
	    connectionFactory.password,
	    connectionFactory.host,
	    connectionFactory.port,
	    connectionFactory.database,
	),
    )
    if err != nil {
	panic(err)
    }

    connectionFactory.connection = connection

    return connection
}

func (connectionFactory *ConnectionFactory) Close() error {
    if connectionFactory.connection == nil {
	return nil
    }

    err := connectionFactory.connection.Close()
    connectionFactory.connection = nil

    return err
}
