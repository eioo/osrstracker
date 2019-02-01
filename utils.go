package main

import (
	"log"
	"os"
	"time"
)

func epochTime() int64 {
	return time.Now().UnixNano() / 1000000
}

func createDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		mkdirErr := os.MkdirAll(path, 700)

		if mkdirErr != nil {
			log.Printf("Failed to create directory \"%s\"", path)
			os.Exit(0)
		}
	}
}

func writeToFile(path string, content *[]byte) {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, createErr = os.Create(path)
		if createErr != nil {
			log.Printf("Failed to create log file \"%s\"", path)
			os.Exit(0)
		}

		defer func() {
			if closeErr := file.Close(); closeErr != nil {
				log.Printf("Failed to close log file \"%s\"", path)
				os.Exit(0)
			}
		}()
	}

	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Printf("Failed to open log file \"%s\"", path)
		os.Exit(0)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("Failed to close log file \"%s\"", path)
			os.Exit(0)
		}
	}()

	// write some text line-by-line to file
	_, err = file.Write(*content)

	// save changes
	err = file.Sync()
	if err != nil {
		log.Printf("Failed to save log file \"%s\"", path)
		os.Exit(0)
	}
}
