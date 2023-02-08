package variables

var ENV string = "prod"

var CONFIGS_DIR string = "../website_configs/"

var CREDENTIALS_PATH string = "../../settings/credentials.json"

var CURRENT_INPUT string = ""

type Item struct {
	Name      string
	Thumbnail string
	Link      string
}

type ItemList struct {
	Website string
	Items   []Item
}
