package packages

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"m3u_consumer/types"
)

func GetDataFromPlaylist(url string) []types.Track {
	var resp *http.Response
	var err error
	req, err := http.NewRequest("GET", url, nil)

	if err == nil {
		req.Header.Set("User-Agent", "Duplecast/1.0 (Android; IPTV)")
		req.Header.Set("Referer", "https://duplecast.com/app")
		req.Header.Set("Origin", "https://duplecast.com")
		req.Header.Set("Connection", "keep-alive")

		resp, err = http.DefaultClient.Do(req)
	}

	if err != nil {
			fmt.Println("Error fetching data:", err)
		}

	if resp != nil {
		defer resp.Body.Close()
	}

	scanner := bufio.NewScanner(resp.Body)

	// Example of this data
	// map[group:Series | Amazon Prime Video group-title:Series | Amazon Prime Video title:Zoo S01E01 tvg-logo:http://gstaticontent.com/images/y7DV6v8YcshWXcAdvItx6Sa5GTD_small.jpg tvg-name:Zoo S01E01]
	var metadata map[string]string
	// Example of this data
	// -1 tvg-name="Zoo S01E01" tvg-logo="http://gstaticontent.com/images/y7DV6v8YcshWXcAdvItx6Sa5GTD_small.jpg" group-title="Series | Amazon Prime Video",Zoo S01E01
	var rawEXTINF string

  metadata_list := make([]types.Track, 0)

	for scanner.Scan() {
		line := scanner.Text()
		
		is_have_prefix := strings.HasPrefix(line, "#EXTINF:")
		// After all #EXTINF: line has url of content like a:
		// #EXTINF:-1 tvg-name="Zoo S01E01" tvg-logo="http://gstaticontent.com/images/y7DV6v8YcshWXcAdvItx6Sa5GTD_small.jpg" group-title="Series | Amazon Prime Video",Zoo S01E01
		// http://f4ntasy.pro:80/series/42F4wrGf/K93Xeh/3367.mp4
		is_content_line := !strings.HasPrefix(line, "#")

		// if line != "" {
		// 	fmt.Println("Line:", line)
		// }

		if is_have_prefix {
			metadata, rawEXTINF = parseEXTINF(line[len("#EXTINF:"):])
			continue
		}

		// fmt.Println("Metadata:", metadata)
		// fmt.Println("Raw EXTINF:", rawEXTINF)

		if len(line) > 0 && metadata != nil && is_content_line {			
				logo := metadata["tvg-logo"]
				title := metadata["title"]
				url := line
				genres := metadata["group"]

				track := types.Track{
					Group:     genres,
					Title:     title,
					Url:       url,
					Logo:      logo,
					RawEXTINF: rawEXTINF, // in API is not used
					IsChannel: isChannel(url, genres),
				}

				metadata_list = append(metadata_list, track) // in API, metadata_list is not used. This is insertion in DB
				metadata = nil // Reset metadata for the next track
		}
	}

	fmt.Println("Total tracks founded:", len(metadata_list))

	return metadata_list
}

func GetPlaylistData(url string) (bufio.Scanner, io.ReadCloser) {
  var resp *http.Response
	var err error
	req, err := http.NewRequest("GET", url, nil)

	if err == nil {
		req.Header.Set("User-Agent", "Duplecast/1.0 (Android; IPTV)")
		req.Header.Set("Referer", "https://duplecast.com/app")
		req.Header.Set("Origin", "https://duplecast.com")
		req.Header.Set("Connection", "keep-alive")

		resp, err = http.DefaultClient.Do(req)
	}

	if err != nil {
    fmt.Println("Error fetching data:", err)
  }

	scanner := bufio.NewScanner(resp.Body)
  return *scanner, resp.Body
}

func parseEXTINF(metadata string) (map[string]string, string) {
	result := make(map[string]string)
	rawEXTINF := metadata
	regex := regexp.MustCompile(`(\w+-*\w*)="([^"]*)"`)
	matches := regex.FindAllStringSubmatch(metadata, -1)

	for _, match := range matches {
		key := match[1]
		value := strings.Trim(match[2], `"`)

		if strings.HasSuffix(key, "-id") {
			result["id"] = value
		} else {
			result[key] = value
		}
	}

	if groupTitle, ok := result["group-title"]; ok {
		parts := strings.SplitN(groupTitle, ",", 2)
		if len(parts) == 2 {
			result["group"] = strings.TrimSpace(parts[0])
		} else {
			result["group"] = strings.TrimSpace(groupTitle)
		}
	}

	lastCommaIndex := strings.LastIndex(metadata, ",")
	if lastCommaIndex != -1 && lastCommaIndex+1 < len(metadata) {
		result["title"] = strings.TrimSpace(metadata[lastCommaIndex+1:])
	}

	return result, rawEXTINF
}

func isChannel(url string, group string) bool {
  if strings.Contains(strings.ToLower(group), "canais") || strings.Contains(strings.ToLower(group), "canal") {
    return true
  }

	return strings.HasSuffix(url, ".ts")
}
