package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// ToPng converts an image to png
func main() {
	// Get JPEG file path (prompt user or use a hardcoded path)
	var filePath string
	fmt.Println("Enter the path to the JPEG file:")
	fmt.Scanln(&filePath)

	// Read JPEG data
	imageData, err := readJpegFile(filePath)
	if err != nil {
		fmt.Println("Error reading JPEG:", err)
		return
	}

	// Convert to GIF
	gifData, err := ToPng(imageData)
	if err != nil {
		fmt.Println("Error converting to GIF:", err)
		return
	}

	// Create output file
	outFile, err := os.Create("output.gif") // Replace with desired output filename
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close() // Close the output file on exit

	// Write GIF data to file
	_, err = outFile.Write(gifData)
	if err != nil {
		fmt.Println("Error writing GIF data:", err)
		return
	}

	fmt.Println("Successfully converted JPEG to GIF!")
}

func readJpegFile(filePath string) ([]byte, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file: %s", filePath)
	}
	defer file.Close() // Close the file on exit

	// Read the file data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file data: %s", filePath)
	}

	return data, nil
}
