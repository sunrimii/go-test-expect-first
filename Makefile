build:
	go build -buildmode=plugin -o=$HOME/expectfirst.so plugin/plugin.go

test:
	cd testdata/src/a; go mod vendor
	go test -v ./...