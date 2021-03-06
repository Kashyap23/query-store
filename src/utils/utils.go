package utils

import (
	"encoding/hex"
	"errors"
	"log"
	"math/rand"
	"os"
)

// Create new dir for the path
func CreateDir(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		err = errors.New("Error creating query-store dir - " + err.Error())
		log.Println(err)
		return err
	}

	log.Println("Created dir - ", path, " successfully")
	return nil
}

// Generate hash
func GenerateRandomHash() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		log.Println("Error generating hash - ", err)
		return "", err
	}

	return hex.EncodeToString(b), nil
}

// Check if if given path exists
func CheckIfPathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		log.Println("Error checking for path - ", err)
		return false
	}

	return true
}
