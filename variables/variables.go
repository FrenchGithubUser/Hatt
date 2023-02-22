package variables

var ENV string = "release"

var CONFIGS_DIR string = "website_configs/"

var CREDENTIALS_PATH string = "settings/credentials.json"

var CURRENT_INPUT string = ""

var RESULTS []ItemList

type Item struct {
	Name      string
	Thumbnail string
	Link      string
	Metadata  map[string]string
}

type ItemList struct {
	Website string
	Items   []Item
}
