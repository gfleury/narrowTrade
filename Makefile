PKGS = $$(go list ./... | grep -v /vendor/)
CMDS = $$(go list ./... | grep /cmd/)

default:
	go build $(CMDS)

test:
	go clean $(PKGS)
	go test $(PKGS) -check.v -coverprofile=coverage.txt -covermode=atomic --tags=unit

integration:
	go clean $(PKGS)
	go test $(PKGS) -check.v -coverprofile=coverage.txt -covermode=atomic --tags=integration

race:
	go clean $(PKGS)
	go test -race $(PKGS) -check.v -coverprofile=coverage.txt -covermode=atomic

profile:
	go clean $(PKGS)
	make
	
clean:
	rm -rf *.prof
	go clean $(PKGS)

lint:
	docker run --rm -it -v ${PWD}:/go/src/github.com/gfleury/narrowTrade -w /go/src/github.com/gfleury/narrowTrade golangci/golangci-lint:latest golangci-lint run
