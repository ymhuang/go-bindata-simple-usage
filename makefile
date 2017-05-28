GOOS=windows
GOARCH=amd64
OUT=sample.exe
BINDATA=HelloWorld.exe
SRC=main.go

.clean-bindata:
	rm -rf res

.build-bindata: .clean-bindata
	mkdir res
	go-bindata -pkg res -o res.go $(BINDATA)
	mv res.go res

all: clean .build-bindata
	go build -o $(OUT) $(SRC)

clean: .clean-bindata
	go clean
	rm -rf $(OUT)
