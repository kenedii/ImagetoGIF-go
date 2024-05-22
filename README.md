# ImagetoGIF-go
Go lang functions to convert an image to static or animated gif

**Convert static image to GIF:**
func ToGif(filePath string, saveFile *bool, saveFileName *string) ([]byte, error)

**Convert folder of image frames to Animated GIF:**
func ToGifA(folder string, saveFile *bool, saveFileName *string) ([]byte, error) 

**Convert image/folder of image frames to GIF through console prompts:**
compile main.go or download the latest release
