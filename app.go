package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"

	packages "m3u_consumer/packages"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
		os.Exit(1)
	}

	URL := os.Getenv("URL")

	fmt.Println("Fetching data from:", URL)
	fmt.Println("Start time:", time.Now().Format(time.RFC1123))

	data := packages.GetDataFromPlaylist(URL)
	dataWithSeriesInfo := packages.AddSeriesInfoIntoTracks(data)
	file, err := os.Create("tracks.txt")

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	for _, t := range dataWithSeriesInfo {
		line := fmt.Sprintf("Group: %s | Title: %s | URL: %s | Logo: %s | IsChannel: %t\n | Season: %s | Episode: %s\n",
			t.Group, t.Title, t.Url, t.Logo, t.IsChannel, t.Season, t.Episode)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Tracks written to tracks.txt successfully.")
	fmt.Println("End time:", time.Now().Format(time.RFC1123))

	// Example of how to get all tracks from a specific series without need to manipulate the title
	// contentFromSeries := packages.GetAllTracksFromSeries(dataWithSeriesInfo, "Van Helsing S01E01", nil)
	// fmt.Println("Total tracks from series 'Van Helsing S01E01':", len(contentFromSeries))
	
	// seriesFile, err := os.Create("series_tracks.txt")
	// if err != nil {
	// 	fmt.Println("Error creating series file:", err)
	// 	return
	// }
	// defer seriesFile.Close()

	// for _, t := range contentFromSeries {
	// 	line := fmt.Sprintf("Group: %s | Title: %s | URL: %s | Logo: %s | IsChannel: %t\n | Season: %s | Episode: %s\n",
	// 		t.Group, t.Title, t.Url, t.Logo, t.IsChannel, t.Season, t.Episode)
	// 	_, err := seriesFile.WriteString(line)
	// 	if err != nil {
	// 		fmt.Println("Error writing to series file:", err)
	// 		return
	// 	}
	// }

	// season := "01" // Example season to filter
	// contentFromSeriesFilteredSeason := packages.GetAllTracksFromSeries(dataWithSeriesInfo, "Van Helsing", &season)
	// fmt.Println("Total tracks from series 'Van Helsing' Season 01:", len(contentFromSeriesFilteredSeason))

	// seriesFileFiltered, err := os.Create("series_tracks_filtered.txt")
	// if err != nil {
	// 	fmt.Println("Error creating filtered series file:", err)
	// 	return
	// }
	// defer seriesFileFiltered.Close()

	// for _, t := range contentFromSeriesFilteredSeason {
	// 	line := fmt.Sprintf("Group: %s | Title: %s | URL: %s | Logo: %s | IsChannel: %t\n | Season: %s | Episode: %s\n",
	// 		t.Group, t.Title, t.Url, t.Logo, t.IsChannel, t.Season, t.Episode)
	// 	_, err := seriesFileFiltered.WriteString(line)
	// 	if err != nil {
	// 		fmt.Println("Error writing to filtered series file:", err)
	// 		return
	// 	}
	// }
}