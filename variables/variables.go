package variables

import (
	"flag"
	"fmt"
	"os"
)

// will only affect the websites used
var ENV string = "release"

// says weather the app is compiled or not
var MODE string = "compiled"

var CONFIGS_DIR string = "website_configs"

var USER_CONFIG_DIR string

var CREDENTIALS_PATH string

var SETTINGS_PATH string

var CUSTOM_LISTS_PATH string

var CURRENT_INPUT string = ""

var SELECTED_CATEGORIES []string

var RESULTS []ItemList

var ARGS []string

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

func InitVariables() {
	flag.Parse()
	fmt.Println("args : ", flag.Args())
	for index, arg := range flag.Args() {
		if index == 0 {
			MODE = arg
		}
	}

	USER_CONFIG_DIR, _ = os.UserConfigDir()
	USER_CONFIG_DIR += "/Hatt"

	CREDENTIALS_PATH = USER_CONFIG_DIR + "/credentials.json"
	SETTINGS_PATH = USER_CONFIG_DIR + "/settings.json"
	CUSTOM_LISTS_PATH = USER_CONFIG_DIR + "/custom_lists.json"

	if ENV == "dev" {
		CONFIGS_DIR += "/dev"
	}

}
