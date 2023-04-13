package variables

type Item struct {
	Name      string
	Thumbnail string
	Link      string
	Metadata  map[string]string
}

type ItemList struct {
	Website               string
	CompatibleDownloaders []CompatibleDownloader
	Items                 []Item
}

var CompatibleDownloaders []CompatibleDownloader

type CompatibleDownloader struct {
	Name string
	Link string
}

type UserSettings struct {
	General General `json:"general"`
}

type General struct {
	ThumbnailsSize    int    `json:"thumbnailsSize"`
	Lang              string `json:"lang"`
	DarkMode          bool   `json:"darkMode"`
	ItemClickedAction string `json:"itemClickedAction"`
	Xxx               bool   `json:"xxx"`
}
