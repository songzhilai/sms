go build -o bin/sendsmssrv_mac64 -ldflags "-s -w" main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/sendsmssrv_linux64 -ldflags "-s -w" main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/sendsmssrv_win64.exe -ldflags "-s -w" main.go
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/sendsmssrv_win32.exe -ldflags "-s -w" main.go