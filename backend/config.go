package main

type Config struct {
	Search     Search   `json:"search"`
	Name       string   `json:"name"`
	Categories []string `json:"categories"`
}

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
