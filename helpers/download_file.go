package helpers

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Returns the next cursor
func DownloadFile(url string, urlValues url.Values, destinationPath string) error {
	resp, err := http.PostForm(url, urlValues)
	if err != nil {
		log.Printf("Failed to get %s", url)
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	f, err := os.Create(destinationPath)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer f.Close()

	f.WriteString(string(body))
	return nil
}
