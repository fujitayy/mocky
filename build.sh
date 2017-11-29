GOOS=darwin GOARCH=amd64 go build
mv mocky mocky.darwin

GOOS=windows GOARCH=amd64 go build
mv mocky.exe mocky.windows.exe

GOOS=linux GOARCH=amd64 go build
mv mocky mocky.linux
