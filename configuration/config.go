package configuration

// This file lists the fields that can be used in the configuration file of the soures

type Config struct {
	SpecificScraper       bool              `json:"specificScraper"`
	SpecificInfo          map[string]string `json:"specificInfo"`
	Search                Search            `json:"search"`
	Login                 Login             `json:"login"`
	Name                  string            `json:"name"`
	Categories            []string          `json:"categories"`
	CompatibleDownloaders []string          `json:"compatibleDownloaders"`
	Xxx                   bool              `json:"xxx"`
}

// ----------------------------------------------

type Search struct {
	Url                        string                     `json:"url"`
	SpaceReplacement           string                     `json:"spaceReplacement"` // character used to replace a space in the user input (usually "-" or "+")
	CategorySpecificAttributes CategorySpecificAttributes `json:"categorySpecificAttributes"`
	Method                     string                     `json:"method"` // POST or GET
	ItemKeys                   ItemKeys                   `json:"itemKeys"`
	PostFields                 PostFields                 `json:"postFields"`
	Encoding                   string                     `json:"encoding"`
}

type CategorySpecificAttributes struct {
	Name   string            `json:"name"`
	Values map[string]string `json:"values"`
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
	OnItemPage      bool   `json:"onItemPage"`
}

type PostFields struct {
	Input   string            `json:"input"`
	Generic map[string]string `json:"generic"`
}

// ----------------------------------------------

// ----------------------------------------------

type Login struct {
	Url                   string            `json:"url"`
	HomeUrl               string            `json:"homeUrl"` // if a token is needed to be sent when logging in (to witness that the user visited the website and isn't just sending login requests)
	CredsFormat           string            `json:"credsFormat"`
	Fields                []string          `json:"fields"`      // fields of the login form
	PageInputs            map[string]string `json:"pageInputs"`  // inputs for the fields of the login form, either this or Fields is used, depending on the login method used
	LoginButton           string            `json:"loginButton"` //only used if PageInputs is used
	GenericFields         map[string]string `json:"genericFields"`
	ServerGeneratedFields []string          `json:"serverGeneratedFields"`
	Headers               map[string]string `json:"headers"`
	AuthMethod            string            `json:"authMethod"`
	TokenLifespan         int               `json:"tokenLifespan"`
	Tokens                []string          `json:"tokens"`
}

// ----------------------------------------------
