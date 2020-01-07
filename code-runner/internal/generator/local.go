package generator

import (
	"context"
	googleGithub "github.com/google/go-github/github")


func (app *GenApp) generateLocalUtils() {
	app.Data = make(map[string][]byte)
	app.Data["makefile"] = app.genMakefile()
	app.Data["bin/_buildImage.sh"] = app.genbuildImage()
	app.Data["bin/_run.sh"] = app.genRun()
	app.Data["bin/_test.sh"] = app.genTest()
	app.Data["bin/_pushImage.sh"] = app.genPushImage()
	app.Data["bin/_pullImage.sh"] = app.genPullImage()
}

func (app *GenApp) genMakefile() []byte {
	makefile:=` 
user=
token=
setup:
	chmod +x bin/*
build:
	chmod +x bin/_buildImage.sh
	bin/_buildImage.sh
test:
	chmod +x bin/_test.sh
	bin/_test.sh
run:
	chmod +x bin/_run.sh
	bin/_run.sh
push:
	chmod +x bin/_pushImage.sh
	bin/_pushImage.sh '$(user)' '$(token)'
pull:
	chmod +x bin/_pullImage.sh
	bin/_pullImage.sh '$(user)' '$(token)'
`
	return []byte(makefile)
}

func (app *GenApp) genbuildImage() []byte {
	cmd:= "`git rev-parse --show-toplevel`"
	script:=
		`#!/usr/bin/env bash
REPO=$(basename `+cmd+`)
docker build . --file Dockerfile --tag $REPO
`
	return []byte(script)
}

func (app *GenApp) genRun() []byte {
	script:=
		`#!/usr/bin/env bash
echo "TBD RUN APP"
`
	return []byte(script)
}

func (app *GenApp) genTest() []byte {
	script:=
		`#!/usr/bin/env bash
echo "TBD TEST APP"
`
	return []byte(script)
}

func (app *GenApp) genPushImage() []byte {
	cmd:="`git rev-parse --show-toplevel`"
	script:=
		`#!/usr/bin/env bash
docker build . --file Dockerfile --tag image
echo "$2" | docker login docker.pkg.github.com -u $1 --password-stdin
REPO=$(basename `+cmd+`)
echo REPO=$REPO
IMAGE_NAME=$(echo "$REPO" | sed -e 's/\//./')
IMAGE_ID="docker.pkg.github.com/$1/$REPO/$1.$IMAGE_NAME"
VERSION="latest"
echo IMAGE_ID:VERSION=$IMAGE_ID:$VERSION
docker tag image "$IMAGE_ID:$VERSION"
docker push $IMAGE_ID:$VERSION
`
	return []byte(script)
}

func (app *GenApp) genPullImage() []byte {
	cmd:="`git rev-parse --show-toplevel`"
	script:=
		`
#!/usr/bin/env bash

echo "$2" | docker login docker.pkg.github.com -u $1 --password-stdin
REPO=$(basename `+cmd+`)
IMAGE_NAME=$(echo "$REPO" | sed -e 's/\//./')
IMAGE_ID="docker.pkg.github.com/$1/$REPO/$IMAGE_NAME"
VERSION="latest"
echo IMAGE_ID:VERSION=$IMAGE_ID:$VERSION
docker pull $IMAGE_ID:$VERSION
`
	return []byte(script)
}

func (app *GenApp)pushLocalUtilsCode(ctx context.Context,user string)error {

	var commitMsg string
	var fileOptions *googleGithub.RepositoryContentFileOptions
	var err error

	for file,content := range app.Data{
		commitMsg="Generating "+file
		fileOptions = BuilFileOptions(commitMsg, user, content)
		_, err = app.CommitFile(ctx, file, fileOptions)
		if err !=nil{
			return err
		}
	}
	return nil
}