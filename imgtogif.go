package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func main() {
	// Get JPEG file path (prompt user or use a hardcoded path)
	var filePath string
	fmt.Println("Enter the path to the JPEG file:")
	fmt.Scanln(&filePath)

	// Convert to GIF
	gifData, err := ToGif(filePath)
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

func ToGif(filePath string) ([]byte, error) {
	var imageData []byte // Initialize with an empty slice of bytes
	var err error

	if filepath.Ext(filePath) == ".jpeg" {
		// Read JPEG data
		imageData, err = readJpegFile(filePath)
		if err != nil {
			fmt.Println("Error reading JPEG:", err)
			return nil, err // Return nil data and the error
		}

	}
	contentType := http.DetectContentType(imageData)

	var img image.Image

	switch contentType {
	case "image/png":
		img, err = png.Decode(bytes.NewReader(imageData))
		if err != nil {
			return nil, errors.Wrap(err, "unable to decode png")
		}
	case "image/jpeg":
		img, err = jpeg.Decode(bytes.NewReader(imageData))
		if err != nil {
			return nil, errors.Wrap(err, "unable to decode jpeg")
		}
	default:
		return nil, fmt.Errorf("unsupported image type: %s", contentType)
	}

	buf := new(bytes.Buffer)
	if err := gif.Encode(buf, img, nil); err != nil {
		return nil, errors.Wrap(err, "unable to encode gif")
	}

	return buf.Bytes(), nil
}
