package config

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Driver   string `default:"sqlite3"`
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Args     string `default:"parseTime=true"`
}

func (c *Database) GetDatabaseDsn() string {
	if c.Driver == "sqlite3" {
		return "file:data/data.db?_fk=1"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", c.User, c.Password, c.Host, c.Port, c.Database, c.Args)
}
