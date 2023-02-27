package variables

// will only affect the websites used
var ENV string = "release"

// says weather the app is compiled or not
var MODE string = "compiled"

var CONFIGS_DIR string = "website_configs/"

var CREDENTIALS_PATH string = "settings/credentials.json"

var CURRENT_INPUT string = ""

var RESULTS []ItemList

var ARGS []string

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
