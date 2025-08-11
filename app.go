package main

import (
	"fmt"
	"time"

	config "m3u_consumer/config"
	controllers "m3u_consumer/controllers"
	packages "m3u_consumer/packages"
	repositories "m3u_consumer/repositories"
)

func main() {
  config.LoadConfig()

	fmt.Println("Fetching data from:", config.URL)
	fmt.Println("Start time:", time.Now().Format(time.RFC1123))

  schemas := []string{
    repositories.Playlists_schema,
  }

  d := packages.NewDatabase()
  db := d.Connect(schemas)

  playlist_controller := controllers.NewPlaylistController(db)
  id, err := playlist_controller.Create(repositories.CreatePlaylistDTO{
    Name:      "My Playlist",
    Pin:       "1234",
    Protected: false,
    Url:       config.URL,
  })

  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Playlist created with ID:", id)
  }
}
