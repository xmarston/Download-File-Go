package testdata

import (
	"testing"
	"download_file"
)

func TestDownloadBytes(t *testing.T) {
	var expectedResult int64 = 32088
	dw := download_file.DownloadFile{}

	bdw, err := dw.Download("https://www.aeed.es/wp-content/uploads/2017/01/google-1.png", "google")
	if err != nil {
		t.Fatal("Expected %d but got %d", expectedResult, bdw)
	}

	if expectedResult != bdw {
		t.Fatal("Expected %d but got %d", expectedResult, bdw)
	}
}

func TestDownloadFilename(t *testing.T) {
	var expectedResult = "google.png"
	dw := download_file.DownloadFile{}

	_, err := dw.Download("https://www.aeed.es/wp-content/uploads/2017/01/google-1.png", "google")
	fn := dw.GetDownloadedFilename()
	if err != nil {
		t.Fatal("Expected %d but got %d", expectedResult, fn)
	}

	if expectedResult != fn {
		t.Fatal("Expected %d but got %d", expectedResult, fn)
	}
}
