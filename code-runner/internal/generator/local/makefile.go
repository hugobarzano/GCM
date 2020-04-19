package local

func GenMakefile() []byte {
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
