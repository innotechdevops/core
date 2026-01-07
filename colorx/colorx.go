package colorx

import (
	"fmt"
	"math"

	"golang.org/x/exp/rand"
)

// Define core hues (red, green, blue, yellow, etc.)
var coreHues = []float64{
	0,   // Red
	30,  // Orange
	60,  // Yellow
	90,  // Yellow-Green
	120, // Green
	150, // Green-Cyan
	180, // Cyan
	210, // Blue-Cyan
	240, // Blue
	270, // Purple
	300, // Purple-Red
	330, // Red-Pink
}

// HSL to RGB conversion
func hslToRgb(h, s, l float64) (int, int, int) {
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60.0, 2)-1))
	m := l - c/2

	var r, g, b float64
	switch {
	case h >= 0 && h < 60:
		r, g, b = c, x, 0
	case h >= 60 && h < 120:
		r, g, b = x, c, 0
	case h >= 120 && h < 180:
		r, g, b = 0, c, x
	case h >= 180 && h < 240:
		r, g, b = 0, x, c
	case h >= 240 && h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}

	// Convert to 0-255 range and return
	r, g, b = (r+m)*255, (g+m)*255, (b+m)*255
	return int(r), int(g), int(b)
}

// Convert RGB to Hexadecimal
func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// Generate a distinct color shade automatically
func generateAutoShade(hue float64) string {
	// Randomly adjust saturation and lightness for variety
	saturation := 0.6 + rand.Float64()*0.4 // Range: 0.6-1.0
	lightness := 0.4 + rand.Float64()*0.2  // Range: 0.4-0.6

	r, g, b := hslToRgb(hue, saturation, lightness)
	return rgbToHex(r, g, b)
}

// Randomly generate a color
func Random() string {
	// Randomly select a hue from the core hues
	hue := coreHues[rand.Intn(len(coreHues))]

	// Automatically generate a distinct shade for the selected hue
	return generateAutoShade(hue)
}

func hue() []float64 {
	var hues []float64
	for i := 0; i < 36; i++ {
		hues = append(hues, float64(i*10))
	}
	return hues
}
