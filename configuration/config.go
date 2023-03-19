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
	Url              string     `json:"url"`
	SpaceReplacement string     `json:"spaceReplacement"` // character used to replace a space in the user input (usually "-" or "+")
	Method           string     `json:"method"`           // POST or GET
	ItemKeys         ItemKeys   `json:"itemKeys"`
	PostFields       PostFields `json:"postFields"`
}

type ItemKeys struct {
	Root      string            `json:"root"`
	Name      string            `json:"name"`
	Thumbnail Thumbnail         `json:"thumbnail"`
	Link      string            `json:"link"`
	Metadata  map[string]string `json:"metadata"`
}

type Thumbnail struct {
	Key             string `json:"key"`
	Attribute       string `json:"attribute"`
	AppendToSiteUrl bool   `json:"appendToSiteUrl"`
}

type PostFields struct {
	Input   string            `json:"input"`
	Generic map[string]string `json:"generic"`
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
