// Package config provides configuration settings for the application.
package config

import (
	"flag"
	"os"
	"sync"
)

// Config holds all configuration for the application
type Config struct {
	Port   int
	Env    string
	API    string // This will be deprecated as we are merging services
	DB     struct {
		DSN string
	}
	Stripe struct {
		Secret string
		Key    string
	}
}

var (
	once     sync.Once
	instance *Config
)

// LoadConfigOnce parses command-line flags and environment variables 
// to populate the Config struct. It ensures this is done only once.
func LoadConfigOnce() *Config {
	once.Do(func() {
		instance = &Config{}
		// The single binary will run on one port
		flag.IntVar(&instance.Port, "port", 4000, "server port to listen on")
		flag.StringVar(&instance.Env, "env", "development", "Environment (development|production)")

		flag.Parse()

		instance.Stripe.Key = os.Getenv("STRIPE_KEY")
		instance.Stripe.Secret = os.Getenv("STRIPE_SECRET")
	})

	return instance
}
