package packages

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetDataFromPlaylist(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env")
	}

	URL := os.Getenv("URL")
	data := GetDataFromPlaylist(URL)

	if len(data) == 0 {
		t.Error("Expected to get some data from the playlist, but got none.")
	}
}

func TestParseEXTINF(t *testing.T) {
  line := `#EXTINF:-1 tvg-name="Zoo S01E01" tvg-logo="http://gstaticontent.com/images/y7DV6v8YcshWXcAdvItx6Sa5GTD_small.jpg" group-title="Series | Amazon Prime Video",Zoo S01E01`
  metadata, rawEXTINF := parseEXTINF(line[len("#EXTINF:"):])

  if metadata == nil {
    t.Error("Expected metadata to be parsed, but got nil.")
  }

  if rawEXTINF == "" {
    t.Error("Expected rawEXTINF to be non-empty, but got empty string.")
  }

  if metadata["tvg-name"] == "" {
    t.Error("Expected 'tvg-name' property in metadata to be non-empty, but got empty or missing.")
  }
  if metadata["group-title"] == "" {
    t.Error("Expected 'group-title' property in metadata to be non-empty, but got empty or missing.")
  }
  if metadata["tvg-logo"] == "" {
    t.Error("Expected 'tvg-logo' property in metadata to be non-empty, but got empty or missing.")
  }
}

func TestIsChannel(t *testing.T) {
  group := "CANAIS | ABERTOS"
  url := "http://xtrpro.cyou:80/621500634/174514897/9577"

  if !isChannel(url, group) {
    t.Error("Expected isChannel to return true for a channel URL, but got false.")
  }
}
