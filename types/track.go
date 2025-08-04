package types

type Track struct {
	Group     string
	Title     string
	Url       string
	Logo      string
	RawEXTINF string
	IsChannel bool
}

type TrackWithSeriesInfo struct {
	Track

	Season  string
	Episode string

	IsSeries   bool
	SeriesName string
}
