del BugReport_*.zip
del eurekalog-server-win32.exe
del eurekalog-server-linux-x64
del eurekalog-server.log

gofmt -w eurekalog-crash-server.go

SET GOARCH=386
SET GOOS=windows
go build -o eurekalog-server-win32.exe eurekalog-crash-server.go

SET GOARCH=amd64
SET GOOS=linux
go build -o eurekalog-server-linux-x64 eurekalog-crash-server.go
