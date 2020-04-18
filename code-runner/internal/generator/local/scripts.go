package local


func GenbuildImage() []byte {
	cmd:= "`git rev-parse --show-toplevel`"
	script:=
		`#!/usr/bin/env bash
REPO=$(basename `+cmd+`)
docker build . --file Dockerfile --tag $REPO
`
	return []byte(script)
}

func GenRun(port string) []byte {
	cmd:= "`git rev-parse --show-toplevel`"
	appPort:=port+":"+port
	script:=
		`#!/usr/bin/env bash
REPO=$(basename `+cmd+`)
docker run -p `+appPort+ ` $REPO
`
	return []byte(script)
}

func GenTest() []byte {
	script:=
		`#!/usr/bin/env bash
echo "TBD TEST APP"
`
	return []byte(script)
}

func GenPushImage() []byte {
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

func GenPullImage() []byte {
	cmd:="`git rev-parse --show-toplevel`"
	script:=
		`
#!/usr/bin/env bash

echo "$2" | docker login docker.pkg.github.com -u $1 --password-stdin
REPO=$(basename `+cmd+`)
IMAGE_NAME=$(echo "$REPO" | sed -e 's/\//./')
IMAGE_ID="docker.pkg.github.com/$1/$REPO/$1.$IMAGE_NAME"
VERSION="latest"
echo IMAGE_ID:VERSION=$IMAGE_ID:$VERSION
docker pull $IMAGE_ID:$VERSION
`
	return []byte(script)
}
