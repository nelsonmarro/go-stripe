// Package models handles data structures for the API.
package models

type Card struct {
	Secret   string `json:"secret"`
	Key      string `json:"key"`
	Currency string `json:"currency"`
}
