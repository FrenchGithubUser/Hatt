package main

type Config struct {
	Search Search `json:"search"`
}

type Search struct {
	Url      string   `json:"url"`
	ItemKeys ItemKeys `json:"itemKeys"`
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
