package model

type AIModel struct {
	ID           uint   `json:"id"`
	PerFilName   string `json:"perfilName"`
	Instructions string `json:"instructions"`
}
