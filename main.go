package main

import (
	"fmt"
	"strconv"

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
		// Get amount time between frames in milliseconds
		var framePeriod int
		fmt.Println("Enter the amount time between frames in milliseconds (press enter for 0)")
		// Scan the user input into a temporary variable
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil || input == "" {
			// Handle error or empty input (set a default value)
			framePeriod = 0
		} else {
			// Convert the input string to int using strconv.Atoi
			f, err := strconv.Atoi(input)
			if err != nil {
				// Handle error (set a default value)
				framePeriod = 0
			} else {
				framePeriod = f / 10
			}
		}

		// Convert to GIF
		_gifData, err := imgtogif.ToGifA(filePath, &saveFiles, &saveFilePath, &framePeriod)
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
