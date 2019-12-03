package tools

import (
	"code-runner/internal/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(url, localPath string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	createDirIfNotExist(filepath.Dir(localPath))
	out, err := os.Create(localPath)

	if err != nil {
		fmt.Printf("error creating folder"+err.Error())
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func createDirIfNotExist(dir string) {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Printf("Error creating folder" + dir)
		}
	}
}

func GetAppLocalPath(app models.App) string {
	return fmt.Sprintf("%v%v/%v/%v.tar.gz",
		os.TempDir(), app.Owner, app.Name, app.Name)
}
