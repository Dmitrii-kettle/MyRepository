package bdconnect

import (
	"MyServer/mlog"
	"database/sql"
	"log"
	"sync"

	_ "modernc.org/sqlite"
)

type DBConnector struct {
	Connector  *sql.DB
	cacheMutex sync.Mutex
}

func GetNewDBConnector() (*DBConnector, error) {
	connector, err := sql.Open("sqlite", "database.db")
	if err = connector.Ping(); err != nil {
		log.Panic("Failed connection to Database")
	}
	connector.SetMaxIdleConns(5)
	mlog.Trace("Database was connected")

	return &DBConnector{
		Connector:  connector,
		cacheMutex: sync.Mutex{},
	}, err
}
