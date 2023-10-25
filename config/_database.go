package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DbConnection interface {
	Conn() *sql.DB
}

type dbConnection struct {
	db  *sql.DB
	cfg *Config
}

func (d *dbConnection) initDb() error {
	var dbConf = d.cfg.DbConfig
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		dbConf.Password,
		dbConf.Name)
	db, err := sql.Open(dbConf.Driver, dataSourceName)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	d.db = db
	return nil
}

// receiver
func (d *dbConnection) Conn() *sql.DB {
	return d.db
}

func NewDbConnection(configParam *Config) (DbConnection, error) {
	conn := &dbConnection{
		cfg: configParam,
	}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
