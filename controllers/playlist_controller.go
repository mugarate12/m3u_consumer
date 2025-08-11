package controllers

import (
	"fmt"
	"io"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	config "m3u_consumer/config"
	packages "m3u_consumer/packages"
	repositories "m3u_consumer/repositories"
)

type playlistController interface {
  Create(playlist repositories.CreatePlaylistDTO) (int, error)
}

type playlistControllerImpl struct {
  db *sqlx.DB
}

func NewPlaylistController(db *sqlx.DB) playlistController {
  return &playlistControllerImpl{db: db}
}


func (c *playlistControllerImpl) Create(playlist repositories.CreatePlaylistDTO) (int, error) {  
  playlists_repository := repositories.NewPlaylistRepository(c.db)
  id, err := playlists_repository.CreatePlaylist(playlist)

  if err != nil {
    return 0, err
  }

  fmt.Println("Fetching data from:", config.URL)
  _, respBody := packages.GetPlaylistData(config.URL)
  fmt.Println("Data fetched successfully.")

  bodyBytes, _ := io.ReadAll(respBody)
  respText := string(bodyBytes)

  err = playlists_repository.UpdatePlaylist(repositories.UpdatePlaylistDTO{
    Id: id,
    Data: &respText,
  })

  return id, err
}
