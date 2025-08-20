package modules

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/kbinani/screenshot"
)

func Screen() ([]string, error) {
	n := screenshot.NumActiveDisplays()
	var files []string

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("screenshot_%d.png", i)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		png.Encode(file, img)
		file.Close()
		files = append(files, filename)
	}
	return files, nil
}
