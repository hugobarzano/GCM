package tools

import "testing"

func TestDownloadFile(t *testing.T) {

	DownloadFile(
		"https://codeload.github.com/hugobarzano/djan/legacy.tar.gz/master",
		"./file.test")
}
