package configuration

type Config struct {
	SpecificScraper       bool              `json:"specificScraper"`
	SpecificInfo          map[string]string `json:"specificInfo"`
	Search                Search            `json:"search"`
	Login                 Login             `json:"login"`
	Name                  string            `json:"name"`
	Categories            []string          `json:"categories"`
	CompatibleDownloaders []string          `json:"compatibleDownloaders"`
}

// ----------------------------------------------

type Search struct {
	Url string `json:"url"`
	// character used to replace a space in the user input (usually "-" or "+")
	SpaceReplacement string   `json:"spaceReplacement"`
	ItemKeys         ItemKeys `json:"itemKeys"`
}

type ItemKeys struct {
	Root      string    `json:"root"`
	Name      string    `json:"name"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Link      string    `json:"link"`
}

type Thumbnail struct {
	Key       string `json:"key"`
	Attribute string `json:"attribute"`
}

// ----------------------------------------------

// ----------------------------------------------

type Login struct {
	Url         string   `json:"url"`
	CredsFormat string   `json:"credsFormat"`
	Fields      []string `json:"fields"`
	AuthMethod  string   `json:"authMethod"`
	Tokens      []string `json:"tokens"`
}

// ----------------------------------------------
