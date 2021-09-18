package model

// FileInfo godoc
type FileInfo struct {
	Size       float64 `json:"size"`
	SizeStr    string  `json:"size_string"`
	Name       string  `json:"name"`
	ModifiedAt string  `json:"modified_at"`
}

// DirInfo godoc
type DirInfo struct {
	Size       float64 `json:"size"`
	SizeStr    string  `json:"size_string"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	IsDownload bool    `json:"is_download"`
	ModifiedAt string  `json:"modified_at"`
}

// Nomor 2
type Relation struct {
	InfluencedBy  []string `json:"influenced-by"`
	Influences 	  []string `json:"influences"`
}
type LanguageHardCoded struct {
	Language string `json:"language"`
	Appeared int32	`json:"appeared"`
	Created	 []string `json:"created"` 
	Functional bool   `json:"functional"`
	ObjectOriented bool `json:"object-oriented"`
	Relation Relation `json:"relation"`


}