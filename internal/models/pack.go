package models

type PackResult struct {
	Package *Package `json:"package"`
	Amount  uint     `json:"amount"`
}
