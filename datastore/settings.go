package datastore

import "database/sql"

type SettingsRepository interface {
	// Connection return the connection of auth database
	Connection() *sql.DB

	// Dismount close the connection of database
	Dismount() error
}
