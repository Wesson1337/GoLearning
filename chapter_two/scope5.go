package main

import (
	"log"
	"os"
)

var cwd string
func init() {
	cwd, err := os.Getwd() // compile error: unused: cwd
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

func init() {
	cwd, err := os.Getwd() // incorrect
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Println(cwd)
}

// correct
func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}