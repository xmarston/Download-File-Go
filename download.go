package download_file

import (
	"net/url"
	"net/http"
	"errors"
	"fmt"
	"strings"
	"os"
	"io"
)

type DownloadFile struct {
	Url      url.URL
	Filename string
}

// Inits the DownloadFile struct in order to set the minimal data to
// allow the download
func (dw *DownloadFile) New(uri, fn string) error {
	parsed, err := url.Parse(uri)
	if err != nil {
		return errors.New(fmt.Sprintf("download_file: %s", err))
	}

	dw.Url = *parsed

	if fn == "" {
		return errors.New("filename could not be empty")
	}
	dw.Filename = fn

	return nil
}

// Download the file from internet and stores it with the name the user inputted
func (dw *DownloadFile) Download() (int64, error) {
	var (
		bdw int64
	)
	resp, err := http.Get(dw.Url.String())

	if err != nil {
		return bdw, errors.New(fmt.Sprintf("download_file get action: %s", err))
	}
	defer resp.Body.Close()
	extension := strings.Split(resp.Header.Get("Content-type"), "/")[1]

	output, err := os.Create(dw.Filename + "." + extension)
	if err != nil {
		return bdw, errors.New(fmt.Sprintf("download_file file creation: %s", err))
	}
	defer output.Close()

	bdw, err = io.Copy(output, resp.Body)
	if err != nil {
		return bdw, errors.New(fmt.Sprintf("download_file copy download file to local file: %s", err))
	}

	return bdw, nil
}
