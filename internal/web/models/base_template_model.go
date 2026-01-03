// Package models holds the app models
package models

type BaseTemplate struct {
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated bool
	API             string
	CSSVersion      string
}
