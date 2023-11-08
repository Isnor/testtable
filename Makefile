
build:
	go fmt ./...
	go vet ./...
	go build ./...

test:
	go test ./...

coverage:
	go test -v ./... -coverpkg=./... -coverprofile=coverage.out

serve-cover: coverage	
	go tool cover -html=coverage.out

serve-doc:
	godoc -http=:6060