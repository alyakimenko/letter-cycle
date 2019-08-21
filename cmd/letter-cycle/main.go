package main

import (
	"math/rand"
	"time"
)

const (
	letterNum = 26

	fullRotation = 360.0
	rotation     = 30.0 // The rotation angle
	step         = .03  // Transparency step

	width  = 200
	height = 200
)

var transperency = 1.0

func init() {
	// Seeds to generate random letters
	rand.Seed(time.Now().UnixNano())
}

func main() {
}

// Returns random character
func getRandomChar() string {
	return string('a' + rand.Intn(letterNum))
}
