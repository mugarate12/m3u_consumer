package repositories

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	config "m3u_consumer/config"
	types "m3u_consumer/types"
)

// DATABASE_URL=postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB

var schema = `
CREATE TABLE IF NOT EXISTS tracks (
  id SERIAL PRIMARY KEY,

  group_name TEXT NOT NULL,
  title TEXT NOT NULL,
  url TEXT NOT NULL,
  logo TEXT,

  season TEXT DEFAULT NULL,
  episode TEXT DEFAULT NULL,

  is_channel BOOLEAN NOT NULL DEFAULT FALSE,
  is_series BOOLEAN NOT NULL DEFAULT FALSE,
  series_name TEXT DEFAULT NULL
)
`

type TrackDTO struct {
  Group      string `db:"group_name"`
  Title      string `db:"title"`
  Url        string `db:"url"`
  Logo       string `db:"logo"`
  Season     string `db:"season"`
  Episode    string `db:"episode"`
  IsChannel  bool   `db:"is_channel"`
  IsSeries   bool   `db:"is_series"`
  SeriesName string `db:"series_name"`
}

func SaveTracksToDatabase(tracks []types.TrackWithSeriesInfo) error {
  db := connectWithDatabase()
  defer db.Close()

  tx := db.MustBegin()

  for _, track := range tracks {
    trackDTO := TrackDTO{
      Group:      track.Group,
      Title:      track.Title,
      Url:        track.Url,
      Logo:       track.Logo,
      Season:     track.Season,
      Episode:    track.Episode,
      IsChannel:  track.IsChannel,
      IsSeries:   track.IsSeries,
      SeriesName: track.SeriesName,
    }

    query := `INSERT INTO tracks (group_name, title, url, logo, season, episode, is_channel, is_series, series_name) 
              VALUES (:group_name, :title, :url, :logo, :season, :episode, :is_channel, :is_series, :series_name)`
    
    if _, err := tx.NamedExec(query, &trackDTO); err != nil {
      log.Println("Error inserting track:", err)
      return err
    }
  }

  if err := tx.Commit(); err != nil {
    log.Println("Error committing transaction:", err)
    return err
  }

  fmt.Println("Tracks saved to database successfully.")
  return nil
}

func connectWithDatabase() *sqlx.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.POSTGRES_HOST, config.POSTGRES_PORT, config.POSTGRES_USER, config.POSTGRES_PASSWORD, config.POSTGRES_DB,
	)

  db, err := sqlx.Connect("postgres", psqlInfo)
  if err != nil {
      log.Fatalln(err)
  }

  db.MustExec(schema)

  return db
}
