package packages

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	config "m3u_consumer/config"
)

type Database interface {
  Connect(schemas []string) *sqlx.DB
}

type databaseImpl struct{}

func NewDatabase() Database {
  return &databaseImpl{}
}

func (d *databaseImpl) Connect(schemas []string) *sqlx.DB {
  fmt.Println("config: ", config.POSTGRES_HOST, config.POSTGRES_PORT, config.POSTGRES_USER, config.POSTGRES_DB)

  psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.POSTGRES_HOST, config.POSTGRES_PORT, config.POSTGRES_USER, config.POSTGRES_PASSWORD, config.POSTGRES_DB,
	)

  db, err := sqlx.Connect("postgres", psqlInfo)
  if err != nil {
      log.Fatalln(err)
  }

  for _, schema := range schemas {
    db.MustExec(schema)
  }

  return db
}
