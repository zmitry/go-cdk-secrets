package utils

import (
	"log"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

// Loads .env file without error
func TryLoad(files ...string) {
	for _, filename := range files {
		err := Load(filename) //nolint
		if err != nil {
			log.Printf("Failed to load %s", err)
		}
	}
}

// Loads .env files relative to current binary
func Load(files ...string) (err error) {
	_, fn, _, _ := runtime.Caller(2)
	for _, filename := range files {
		p := filename
		if !path.IsAbs(filename) {
			p = path.Join(path.Dir(fn), filename)
		}
		err = godotenv.Load(p)
		if err != nil {
			return
		}
	}
	return
}
