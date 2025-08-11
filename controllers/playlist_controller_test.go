package controllers

import (
	"testing"

	config "m3u_consumer/config"
	packages "m3u_consumer/packages"
	repositories "m3u_consumer/repositories"
)

func TestCreatePlaylist(t *testing.T) {
  config.LoadConfig("../.env")

  database_instance := packages.NewDatabase()
  db := database_instance.Connect([]string{
    repositories.Playlists_schema,
  })
  controller := NewPlaylistController(db)

  playlist := repositories.CreatePlaylistDTO{
    Name:      "Test Playlist",
    Pin:       "1234",
    Protected: false,
    Url:       config.URL,
  }

  id, err := controller.Create(playlist)
  if err != nil {
    t.Fatalf("Failed to create playlist: %v", err)
  }

  if id <= 0 {
    t.Errorf("Expected a valid ID, got %d", id)
  }
}
