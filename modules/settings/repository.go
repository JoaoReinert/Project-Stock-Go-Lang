package settings

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/entities"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type settingsRepository struct {
	config     entities.Config
	connection *sql.DB
}

func NewSettingsRepository(config entities.Config) datastore.SettingsRepository {
	settings := &settingsRepository{
		config:     config,
		connection: nil,
	}

	settings.openConnection()
	return settings
}

func (s *settingsRepository) Connection() *sql.DB {
	return s.connection
}

func (s *settingsRepository) Dismount() error {
	return s.connection.Close()
}

func (s *settingsRepository) openConnection() {

	log.Printf(s.config.Database.User)
	log.Printf(s.config.Database.Password)
	log.Printf(s.config.Database.Host)
	log.Printf(s.config.Database.Port)
	log.Printf(s.config.Database.Database)

	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		s.config.Database.User,
		s.config.Database.Password,
		s.config.Database.Host,
		s.config.Database.Port,
		s.config.Database.Database,
	)

	conn, err := sql.Open("mysql", source)
	if err != nil {
		log.Printf("Error openning the database conn [%s] | [%v]", source, err)
		panic(err)
	}

	err = conn.PingContext(context.Background())
	if err != nil {
		log.Printf("Error openning the database conn [%s] | [%v]", source, err)
		panic(err)
	}

	conn.SetMaxOpenConns(20)
	conn.SetConnMaxLifetime(20 * time.Minute)
	conn.SetConnMaxIdleTime(20 * time.Minute)

	s.connection = conn
	return
}
