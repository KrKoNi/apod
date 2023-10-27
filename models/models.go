package models

type ApodModel struct {
	Copyright      string `json:"copyright" db:"copyright"`
	Date           string `json:"date" db:"date"`
	Explanation    string `json:"explanation" db:"explanation"`
	HdUrl          string `json:"hdurl" db:"hdurl"`
	MediaType      string `json:"media_type" db:"media_type"`
	ServiceVersion string `json:"service_version" db:"service_version"`
	Title          string `json:"title" db:"title"`
	Url            string `json:"url" db:"url"`
	Content        []byte `db:"content"`
}
