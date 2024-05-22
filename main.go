package main

import (
	"fmt"

	"imgtogif"
)

func main() {
	var saveFiles bool = true

	// Create static or animated gif?
	var isAnimated string
	fmt.Println("Enter 's' for static or 'a' for animated:")
	fmt.Scanln(&isAnimated)

	// Get JPEG file path (prompt user or use a hardcoded path)
	var filePath string
	fmt.Println("Enter the path to the image file or folder name:")
	fmt.Scanln(&filePath)

	// Get name of GIF file to save
	var saveFilePath string
	fmt.Println("Enter the file name of new GIF file:")
	fmt.Scanln(&saveFilePath)

	// Convert to Animated GIF
	if isAnimated == "a" {
		// Convert to GIF
		_gifData, err := imgtogif.ToGifA(filePath, &saveFiles, &saveFilePath)
		if err != nil {
			fmt.Println("Error converting to GIF:", err)
			return
		}
		_ = _gifData // Ignore the variable to avoid the "declared and not used" error
	}

	// Convert to Static GIF
	if isAnimated == "s" {
		// Convert to GIF
		_gifData, err := imgtogif.ToGif(filePath, &saveFiles, &saveFilePath)
		if err != nil {
			fmt.Println("Error converting to GIF:", err)
			return
		}
		_ = _gifData // Ignore the variable to avoid the "declared and not used" error
	}

	fmt.Println("Successfully converted image(s) to GIF!")
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
