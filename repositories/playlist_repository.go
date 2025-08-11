package repositories

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Playlists_schema = `
CREATE TABLE IF NOT EXISTS playlists (
  id SERIAL PRIMARY KEY,

  name TEXT NOT NULL,
  pin TEXT NOT NULL,
  protected BOOLEAN NOT NULL DEFAULT FALSE,
  url TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'active',
  data TEXT NULL
)
`

type PlaylistDTO struct {
	Name      string `db:"name"`
	Pin       string `db:"pin"`
	Protected bool   `db:"protected"`
	Url       string `db:"url"`
	Status    string `db:"status"`
	Data      string `db:"data"`
}

type CreatePlaylistDTO struct {
  Name      string `db:"name"`
  Pin       string `db:"pin"`
  Protected bool   `db:"protected"`
  Url       string `db:"url"`
}

type UpdatePlaylistDTO struct {
  Id        int     `db:"id"`
  Name      *string `db:"name"`
  Pin       *string `db:"pin"`
  Protected *bool   `db:"protected"`
  Url       *string `db:"url"`
  Status    *string `db:"status"`
  Data      *string `db:"data"`
}

type playlistRepository interface {
  CreatePlaylist(playlist CreatePlaylistDTO) (int, error)
  UpdatePlaylist(playlist UpdatePlaylistDTO) error
  // GetPlaylistById(id int) (*PlaylistDTO, error)
}

type playlistRepositoryImpl struct {
  db *sqlx.DB
}

func NewPlaylistRepository(db *sqlx.DB) playlistRepository {
  return &playlistRepositoryImpl{db: db}
}

func (r *playlistRepositoryImpl) CreatePlaylist(playlist CreatePlaylistDTO) (int, error) {
  query := `
    INSERT INTO playlists (name, pin, protected, url)
    VALUES (:name, :pin, :protected, :url)
    RETURNING id
  `
  var id int
  rows, err := r.db.NamedQuery(query, playlist)
  if err != nil {
    return 0, err
  }

  defer rows.Close()
  
  if rows.Next() {
    if err := rows.Scan(&id); err != nil {
      return 0, err
    }
    return id, nil
  }
  
  return 0, nil
}

func (r *playlistRepositoryImpl) UpdatePlaylist(playlist UpdatePlaylistDTO) error {
  query := `
    UPDATE playlists
    SET name = COALESCE(:name, name),
        pin = COALESCE(:pin, pin),
        protected = COALESCE(:protected, protected),
        url = COALESCE(:url, url),
        status = COALESCE(:status, status),
        data = COALESCE(:data, data)
    WHERE id = :id
  `

  _, err := r.db.NamedExec(query, playlist)
  if err != nil {
    return err
  }

  return nil
}
