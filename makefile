all:
	rm -rf res
	rm -rf main.exe
	mkdir res
	go-bindata -pkg res -o res.go HelloWorld.exe
	mv res.go res
	GOOS=windows GOARCH=amd64 go build main.go


