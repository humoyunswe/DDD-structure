package common

// NameResponse is a reusable multilingual name DTO (same idea as szpt_new).
type NameResponse struct {
	Ru string `json:"ru"`
	En string `json:"en"`
	Kz string `json:"kz"`
}
