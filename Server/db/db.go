package db

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
)

type Connection interface {
	Close()
	GetDatabase() *mgo.Database
}

type conn struct {
	session  *mgo.Session
	database *mgo.Database
}

func NewConnection(cfg Config) (Connection, error) {
	fmt.Printf("Database URL: %s\n", cfg.GetDsn())
	session, err := mgo.Dial(cfg.GetDsn())
	if err != nil {
		return nil, err
	}
	return &conn{
		session:  session,
		database: session.DB(cfg.GetDbName()),
	}, nil
}

func (c *conn) Close() {
	c.session.Close()
}

func (c *conn) GetDatabase() *mgo.Database {
	return c.database
}
