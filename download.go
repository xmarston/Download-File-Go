package download_file

import (
	"net/http"
	"errors"
	"fmt"
	"strings"
	"os"
	"io"
	"time"
)

type DownloadFile struct {
	filename string
}

// Download the file from internet and stores it with the name the user inputted
func (dw *DownloadFile) Download(url, fn string) (int64, error) {
	var (
		bdw int64
	)

	tr := &http.Transport{
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 20,
	}
	netClient := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(20 * time.Second),
	}

	resp, err := netClient.Get(url)

	if err != nil {
		return bdw, errors.New(fmt.Sprintf("download_file get action: %s", err))
	}
	defer resp.Body.Close()

	extension := strings.Split(resp.Header.Get("Content-type"), "/")[1]

	output, err := os.Create(fn + "." + extension)
	if err != nil {
		return bdw, errors.New(fmt.Sprintf("download_file file creation: %s", err))
	}
	defer output.Close()

	bdw, err = io.Copy(output, resp.Body)
	if err != nil {
		return bdw, errors.New(fmt.Sprintf("download_file copy download file to local file: %s", err))
	}

	dw.filename = fn + "." + extension

	return bdw, nil
}

// Gets the filename with the right file extension
func (dw *DownloadFile) GetDownloadedFilename() string {
	return dw.filename
}
