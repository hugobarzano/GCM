package generator

import (
	"code-runner/internal/generator/ci"
	"context"
	googleGithub "github.com/google/go-github/v31/github"
)

func (app *GenApp) generateCI() {
	app.CI = make(map[string][]byte)
	//app.CI[".github/workflows/fastci.yml"]=ci.FastImageBuilder()
	app.CI[".github/workflows/ci.yml"] = ci.ImageBuilder()
}

func (app *GenApp) DisableCI() {
	app.CI = make(map[string][]byte)
	app.CI[".github/workflows/ci.yml"] = []byte{0, 0, 0, 0, 0}
}

func (app *GenApp) PushCI(ctx context.Context, user, mail string) error {
	var commitMsg string
	var fileOptions *googleGithub.RepositoryContentFileOptions
	var err error
	for file, content := range app.CI {
		commitMsg = "Generating CI workflow action " + file
		fileOptions = BuildFileOptions(commitMsg, user, mail, content)
		_, err = app.CommitFile(ctx, file, fileOptions)
		if err != nil {
			return err
		}
	}
	return nil
}
