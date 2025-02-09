build:
	GOOS=windows GOARCH=amd64 go build -o build/chroma.exe
	GOOS=linux GOARCH=amd64 go build -o build/chroma
