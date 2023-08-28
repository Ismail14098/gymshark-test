package errors

import "fmt"

// ErrSystem
const ErrSystem = "GYMSHARK"

// Error structure
type ArgError struct {
	System           string `json:"system"`
	Status           int    `json:"status"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

// Error to formatted string
func (e *ArgError) Error() string {
	return fmt.Sprintf("%d %s", e.Status, e.DeveloperMessage)
}

// Set developer message and return
func (e *ArgError) SetDevMessage(developMessage string) *ArgError {
	e.DeveloperMessage = developMessage
	return e
}
