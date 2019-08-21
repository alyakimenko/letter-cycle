package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	svg "github.com/ajstarks/svgo"
)

const (
	letterNum = 26

	fullRotation = 360.0
	rotation     = 30.0 // The rotation angle
	step         = .03  // Transparency step

	width  = 200
	height = 200
)

var (
	transperency = 1.0
	canvas       *svg.SVG // Main scetch area
)

func init() {
	// Seeds to generate random letters
	rand.Seed(time.Now().UnixNano())
}

func main() {
	scetchWidth := width / 2
	scetchHeight := height / 2

	initCanvas(scetchWidth, scetchHeight)
}

// Inits basic canvas, and define its draw area
func initCanvas(scetchWidth, scetchHeight int) {
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	canvas.Gstyle("font-family: serif; fill: white; font-size: 72pt")
	canvas.Gtransform(fmt.Sprintf("translate(%d, %d)", scetchWidth, scetchHeight))
}

// Returns random character
func getRandomChar() string {
	return string('a' + rand.Intn(letterNum))
}
