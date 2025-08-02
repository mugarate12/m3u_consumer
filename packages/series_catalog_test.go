package packages

import (
	"testing"

	tracks_types "m3u_consumer/types"
)

func TestGetNameOfSeries(t *testing.T) {
  title := "Porto dos Milagres S02E02"
  seriesName := GetNameOfSeries(title)
  expected := "Porto dos Milagres"

  if seriesName != expected {
    t.Errorf("expected '%s', got '%s'", expected, seriesName)
  }
}

func TestGetSeriesInfoFromTitle(t *testing.T) {
  title := "Porto dos Milagres S02E02"
  seriesInfo, ok := getSeriesInfoFromTitle(title)

  if !ok {
    t.Error("expected to find series info, but got false")
  }

  if seriesInfo.Season != "02" || seriesInfo.Episode != "02" {
    t.Errorf("expected season '02' and episode '02', got season '%s' and episode '%s'", seriesInfo.Season, seriesInfo.Episode)
  }
}

func TestAddSeriesInfoIntoTracks(t *testing.T) {
  tracks := []tracks_types.Track{
    {Title: "Porto dos Milagres S01E01", IsChannel: false, Group: "Séries | Novelas", Url: "http://xtrpro.cyou:80/series/621500634/174514897/2125836.mp4", Logo: "http://gstaticontent.com/images/gL9UEJuQCCHk9Fkobpm3kRKpfHv_big.jpg", RawEXTINF: ""},
    {Title: "Porto dos Milagres S01E02", IsChannel: false, Group: "Séries | Novelas", Url: "http://xtrpro.cyou:80/series/621500634/174514897/2125837.mp4", Logo: "http://gstaticontent.com/images/gL9UEJuQCCHk9Fkobpm3kRKpfHv_big.jpg", RawEXTINF: ""},
    {Title: "Terra Prometida (1975)", IsChannel: false, Group: "Filmes | Drama ", Url: "http://xtrpro.cyou:80/movie/621500634/174514897/1740620.mp4", Logo: "https://image.tmdb.org/t/p/w600_and_h900_bestv2/uwMs9axSBS0opiO7Ih2J3AjUvMv.jpg", RawEXTINF: ""},
    {Title: "TV Senado", IsChannel: true, Group: "CANAIS | ABERTOS", Url: "http://xtrpro.cyou:80/621500634/174514897/9603", Logo: "http://www.tmsimg.com/assets/s105047_ll_h3_aa.png", RawEXTINF: ""},
  }

  expected := []TrackWithSeriesInfo{
    {Track: tracks[0], Season: "01", Episode: "01"},
    {Track: tracks[1], Season: "01", Episode: "02"},
    {Track: tracks[2], Season: "", Episode: ""},
    {Track: tracks[3], Season: "", Episode: ""},
  }

  result := AddSeriesInfoIntoTracks(tracks)

  if len(result) != len(expected) {
    t.Fatalf("expected %d tracks, got %d", len(expected), len(result))
  }

  for i, track := range result {
    if track.Season != expected[i].Season || track.Episode != expected[i].Episode {
      t.Errorf("track %d mismatch: expected (%s, %s), got (%s, %s)", i, expected[i].Season, expected[i].Episode, track.Season, track.Episode)
    }
  }
}

func TestGetAllTracksFromSeries(t *testing.T) {
  tracks := []tracks_types.Track{
    {Title: "Porto dos Milagres S01E01", IsChannel: false, Group: "Séries | Novelas", Url: "http://xtrpro.cyou:80/series/621500634/174514897/2125836.mp4", Logo: "http://gstaticontent.com/images/gL9UEJuQCCHk9Fkobpm3kRKpfHv_big.jpg", RawEXTINF: ""},
    {Title: "Porto dos Milagres S02E02", IsChannel: false, Group: "Séries | Novelas", Url: "http://xtrpro.cyou:80/series/621500634/174514897/2125837.mp4", Logo: "http://gstaticontent.com/images/gL9UEJuQCCHk9Fkobpm3kRKpfHv_big.jpg", RawEXTINF: ""},
    {Title: "Terra Prometida (1975)", IsChannel: false, Group: "Filmes | Drama ", Url: "http://xtrpro.cyou:80/movie/621500634/174514897/1740620.mp4", Logo: "https://image.tmdb.org/t/p/w600_and_h900_bestv2/uwMs9axSBS0opiO7Ih2J3AjUvMv.jpg", RawEXTINF: ""},
    {Title: "TV Senado", IsChannel: true, Group: "CANAIS | ABERTOS", Url: "http://xtrpro.cyou:80/621500634/174514897/9603", Logo: "http://www.tmsimg.com/assets/s105047_ll_h3_aa.png", RawEXTINF: ""},
  }

  tracksWithSeries := []TrackWithSeriesInfo{
    {Track: tracks[0], Season: "01", Episode: "01"},
    {Track: tracks[1], Season: "02", Episode: "02"},
    {Track: tracks[2], Season: "", Episode: ""},
    {Track: tracks[3], Season: "", Episode: ""},
  }

  series := GetAllTracksFromSeries(tracksWithSeries, "Porto dos Milagres S01E01", nil)

  if len(series) != 2 {
    t.Fatalf("expected 2 tracks, got %d", len(series))
  }

  season := "01"
  seriesFilteredBySeason := GetAllTracksFromSeries(tracksWithSeries, "Porto dos Milagres S01E01", &season)

  if len(seriesFilteredBySeason) != 1 {
    t.Log("Series filtered by season:", season)
    t.Log("Series tracks:", seriesFilteredBySeason)
    t.Fatalf("expected 1 track for season %s, got %d", season, len(seriesFilteredBySeason))
  }
}
