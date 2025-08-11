package packages

import (
	"testing"

	config "m3u_consumer/config"
)


func TestConnect(t *testing.T) {
  config.LoadConfig("../.env")

  d := NewDatabase()
  db := d.Connect([]string{"CREATE TABLE IF NOT EXISTS test (id SERIAL PRIMARY KEY, name TEXT)"})

  if db == nil {
    t.Error("Expected a valid database connection, but got nil.")
    return;
  }

  // Clean up the test table
  _, err := db.Exec("DROP TABLE IF EXISTS test")
  if err != nil {
    t.Errorf("Failed to drop test table: %v", err)
  }
}
