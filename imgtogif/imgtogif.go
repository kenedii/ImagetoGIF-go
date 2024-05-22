package imgtogif

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/soniakeys/quant/median"
)

func ToGif(filePath string, saveFile *bool, saveFileName *string) ([]byte, error) { // Convert a still JPG or PNG file to GIF
	if saveFileName == nil {
		defaultName := "output"
		saveFileName = &defaultName
	}
	if saveFile == nil {
		defaultSaveValue := true
		saveFile = &defaultSaveValue
	}
	var imageData []byte // Initialize with an empty slice of bytes
	var err error
	// Read the image data
	imageData, err = readImgFile(filePath)
	if err != nil {
		fmt.Println("Error reading JPEG:", err)
		return nil, err // Return nil data and the error
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
	if *saveFile == true {
		// Create output file
		outFile, err := os.Create(*saveFileName + ".gif") // Replace with desired output filename
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return nil, err
		}
		defer outFile.Close() // Close the output file on exit

		// Write GIF data to file
		_, err = outFile.Write(buf.Bytes())
		if err != nil {
			fmt.Println("Error writing GIF data:", err)
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func ToGifA(folder string, saveFile *bool, saveFileName *string) ([]byte, error) {
	if saveFileName == nil {
		defaultName := "output"
		saveFileName = &defaultName
	}
	if saveFile == nil {
		defaultSaveValue := true
		saveFile = &defaultSaveValue
	}

	files, err := ioutil.ReadDir(folder) // Get the list of files in the folder
	if err != nil {
		return nil, errors.Wrap(err, "unable to read folder")
	}

	anim := gif.GIF{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(folder, file.Name())
		imageData, err := readImgFile(filePath)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to read image file: %s", filePath)
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

		bounds := img.Bounds()
		// Use the median cut quantization to preserve colors
		quantizer := median.Quantizer(256)
		palette := quantizer.Quantize(make(color.Palette, 0, 256), img)
		pm := image.NewPaletted(bounds, palette)
		draw.Draw(pm, bounds, img, image.Point{}, draw.Src)
		anim.Image = append(anim.Image, pm)
		anim.Delay = append(anim.Delay, 0) // 0 delay for instant transition between frames
	}

	buf := new(bytes.Buffer)
	if err := gif.EncodeAll(buf, &anim); err != nil {
		return nil, errors.Wrap(err, "unable to encode gif")
	}

	if *saveFile { // If the save file flag is set to true, save the GIF data to a file
		// Create output file
		outFile, err := os.Create(*saveFileName + ".gif") // Replace with desired output filename
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return nil, err
		}
		defer outFile.Close() // Close the output file on exit

		// Write GIF data to file
		_, err = outFile.Write(buf.Bytes())
		if err != nil {
			fmt.Println("Error writing GIF data:", err)
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func readImgFile(filePath string) ([]byte, error) { // Read the Image file data
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
