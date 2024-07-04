# ImagetoGIF-go
Go lang functions to convert an image to static or animated gif

**Convert static image to GIF:**
func ToGif(filePath string, saveFile *bool, saveFileName *string) ([]byte, error)

**Convert folder of image frames to Animated GIF:**
func ToGifA(folder string, saveFile *bool, saveFileName *string, framePeriod *int) ([]byte, error) 

folder: folder of image frames of the gif in order
saveFile: whether to save the file or not
saveFileName: name of destination gif file
framePeriod: number of milliseconds between frames

**Convert image/folder of image frames to GIF through console prompts:**
compile main.go or download the latest release
