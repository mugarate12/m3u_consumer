package repositories

import (
	"testing"

	config "m3u_consumer/config"
	packages "m3u_consumer/packages"
)

func TestCreatePlaylist(t *testing.T) {
  config.LoadConfig("../.env")

  database_instance := packages.NewDatabase()
  db := database_instance.Connect([]string{
    Playlists_schema,
  })
  repo := NewPlaylistRepository(db)

  playlist := CreatePlaylistDTO{
    Name:      "Test Playlist",
    Pin:       "1234",
    Protected: false,
    Url:       "http://example.com/playlist.m3u",
  }

  id, err := repo.CreatePlaylist(playlist)
  if err != nil {
    t.Fatalf("Failed to create playlist: %v", err)
  }

  if id <= 0 {
    t.Errorf("Expected a valid ID, got %d", id)
  }
}

func TestUpdatePlaylist(t *testing.T) {
  config.LoadConfig("../.env")

  database_instance := packages.NewDatabase()
  db := database_instance.Connect([]string{
    Playlists_schema,
  })
  repo := NewPlaylistRepository(db)

  // First, create a playlist to update
  playlist := CreatePlaylistDTO{
    Name:      "Test Playlist",
    Pin:       "1234",
    Protected: false,
    Url:       "http://example.com/playlist.m3u",
  }

  id, err := repo.CreatePlaylist(playlist)
  if err != nil {
    t.Fatalf("Failed to create playlist: %v", err)
  }

  // Now, update the created playlist
  newName := "Updated Playlist"
  newStatus := "inactive"
  updateData := UpdatePlaylistDTO{
    Id:     id,
    Name:   &newName,
    Status: &newStatus,
  }

  err = repo.UpdatePlaylist(updateData)
  if err != nil {
    t.Fatalf("Failed to update playlist: %v", err)
  }
}
