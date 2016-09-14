package conf

import (
	"database/sql"
)

type Config struct {
	Tkpcore	TkpCoreConfig
}

type TkpCoreConfig struct {
	SlaveDB	string
}

type DBBundle struct {
	Core	*sql.DB
}