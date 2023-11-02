package variables

import (
	"flag"
	"os"
)

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

var CURRENT_VERSION string = "0.3.4"

var CURRENT_USER_SETTINGS UserSettings

func InitVariables() {
	flag.Parse()
	// fmt.Println("args : ", flag.Args())
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
}
