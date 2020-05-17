setup:
	chmod +x bin/*
build:
	env GOOS=darwin GOARCH=amd64 go build -o code-runner.darwin main.go;
	env GOOS=linux GOARCH=amd64 go build -o code-runner.linux main.go;
clean:
	rm code-runner.darwin && rm code-runner.linux;