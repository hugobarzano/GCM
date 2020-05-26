package local

func GenbuildImage() []byte {
	cmd := "`git rev-parse --show-toplevel`"
	script :=
		`#!/usr/bin/env bash
REPO=$(basename ` + cmd + `)
docker build . --file Dockerfile --tag $REPO
`
	return []byte(script)
}

func GenRun(port string) []byte {
	cmd := "`git rev-parse --show-toplevel`"
	appPort := port + ":" + port
	script :=
		`#!/usr/bin/env bash
REPO=$(basename ` + cmd + `)
docker run -p ` + appPort + ` $REPO
`
	return []byte(script)
}

func GenTest() []byte {
	script :=
		`#!/usr/bin/env bash
echo "TBD TEST APP"
`
	return []byte(script)
}

func GenPullImage() []byte {
	cmd := "`git rev-parse --show-toplevel`"
	script :=
		`
#!/usr/bin/env bash

echo "$2" | docker login docker.pkg.github.com -u $1 --password-stdin
REPO=$(basename ` + cmd + `)
IMAGE_NAME=$(echo "$REPO" | sed -e 's/\//./')
IMAGE_ID="docker.pkg.github.com/$1/$REPO/$1.$IMAGE_NAME"
VERSION="latest"
echo IMAGE_ID:VERSION=$IMAGE_ID:$VERSION
docker pull $IMAGE_ID:$VERSION
`
	return []byte(script)
}

func GenPushImage() []byte {
	cmd := "`git rev-parse --show-toplevel`"
	version := "`get_latest_release $1/$REPO`"
	minor := "`echo $CURRENT_VERSION | awk -F. '{$NF+=1; OFS=\".\"; print $3}'`"
	mid := "`echo $CURRENT_VERSION | awk -F. '{print $2}'`"
	mayor := "`echo $CURRENT_VERSION | awk -F. '{print $1}'`"
	script :=
		`
#!/usr/bin/env bash
docker build . --file Dockerfile --tag image
echo "$2" | docker login docker.pkg.github.com -u $1 --password-stdin
REPO=$(basename ` + cmd + `)
echo REPO=$REPO
IMAGE_NAME=$(echo "$REPO" | sed -e 's/\//./')
IMAGE_ID="docker.pkg.github.com/$1/$REPO/$1.$IMAGE_NAME"
git fetch

get_latest_release() {
  curl --silent "https://api.github.com/repos/$1/releases/latest" | 
    grep '"tag_name":' |                                            
    sed -E 's/.*"([^"]+)".*/\1/'                                    
}

CURRENT_VERSION=` + version + `
VERSION_MINOR=` + minor + `
VERSION_MID=` + mid + `
VERSION_MAYOR=` + mayor + `
[ -z "$VERSION_MINOR" ] && VERSION_MINOR="0"
[ -z "$VERSION_MID" ] && VERSION_MID="0"
[ -z "$VERSION_MAYOR" ] && VERSION_MAYOR="0"
VERSION="$VERSION_MAYOR.$VERSION_MID.$VERSION_MINOR"
echo VERSION=$VERSION

generate_post_data()
{
cat <<EOF
{
"tag_name": "$VERSION",
"target_commitish": "",
"name": "$VERSION",
"body": "Release to deploy",
"draft": false,
"prerelease": false
}
EOF
}
curl --data "$(generate_post_data)" "https://api.github.com/repos/$1/$REPO/releases?access_token=$2"
echo IMAGE_ID:VERSION=$IMAGE_ID:$VERSION
docker tag image "$IMAGE_ID:$VERSION"
docker push $IMAGE_ID:$VERSION
`
	return []byte(script)
}
