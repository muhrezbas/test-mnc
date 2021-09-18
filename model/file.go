package model

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