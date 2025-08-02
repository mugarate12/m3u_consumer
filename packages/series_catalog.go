package packages

import (
	"regexp"
	"strings"

	tracks_types "m3u_consumer/types"
)

type SeriesInfo struct {
	Season  string
	Episode string
}

func AddSeriesInfoIntoTracks(tracks []tracks_types.Track) []tracks_types.TrackWithSeriesInfo {
	var seriesInfoList []tracks_types.TrackWithSeriesInfo
	
	for _, track := range tracks {
		if track.IsChannel {
			seriesInfoList = append(seriesInfoList, tracks_types.TrackWithSeriesInfo{
				Track:   track,
				Season:  "",
				Episode: "",
			})
			continue
		}

		seriesInfo, ok := getSeriesInfoFromTitle(track.Title)

		if !ok {
			seriesInfoList = append(seriesInfoList, tracks_types.TrackWithSeriesInfo{
				Track:   track,
				Season:  "",
				Episode: "",
			})
			continue
		}

		if seriesInfo.Season != "" && seriesInfo.Episode != "" {
			seriesInfoList = append(seriesInfoList, tracks_types.TrackWithSeriesInfo{
				Track:   track,
				Season:  seriesInfo.Season,
				Episode: seriesInfo.Episode,
			})
		} else {
			seriesInfoList = append(seriesInfoList, tracks_types.TrackWithSeriesInfo{
				Track:   track,
				Season:  "",
				Episode: "",
			})
		}
	}

	return seriesInfoList
}

func GetAllTracksFromSeries (tracks []tracks_types.TrackWithSeriesInfo, title string, season *string) []tracks_types.TrackWithSeriesInfo {
	var tracksFiltered []tracks_types.TrackWithSeriesInfo
	titleWithoutSeason := GetNameOfSeries(title)

	// its can be a LIKE SQL query or a simple SQL query if our includes:
	// isSeries (to indicate that we are looking for series)
	// seriesName (to indicate the name of the series, same that content of titleWithoutSeason above)
	// AND with season info (in column Season) can be possible used to filter by season
	for _, track := range tracks {
		if track.IsChannel {
			continue
		}

		isHaveTitleContent := strings.Contains(strings.ToLower(track.Title), strings.ToLower(titleWithoutSeason))
		if isHaveTitleContent && track.Season != "" && track.Episode != "" {
			tracksFiltered = append(tracksFiltered, track)
		}
	}

	if season != nil {
		var tracksFilteredBySeason []tracks_types.TrackWithSeriesInfo
		for _, track := range tracksFiltered {
			if track.Season == *season {
				tracksFilteredBySeason = append(tracksFilteredBySeason, track)
			}
		}
		
		return tracksFilteredBySeason
	}

	return tracksFiltered
}

// Return the series name without season and episode info
func GetNameOfSeries(title string) string {
	re := regexp.MustCompile(`^(.*?)\s*S\d{1,2}E\d{1,3}`)
	matches := re.FindStringSubmatch(title)

	if len(matches) > 0 {
		return strings.TrimSpace(matches[1])
	}

	return title
}

func getSeriesInfoFromTitle(title string) (SeriesInfo, bool)  {
	re := regexp.MustCompile(`(?i)S(\d{1,2})E(\d{1,3})`)
	matches := re.FindStringSubmatch(title)

	if len(matches) == 3 {
		season := matches[1]
		episode := matches[2]

		return SeriesInfo{
			Season:  season,
			Episode: episode,
		}, true
	}

	return SeriesInfo{}, false
}
