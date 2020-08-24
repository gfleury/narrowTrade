PKGS = $$(go list ./... | grep -v /vendor/)
CMDS = $$(go list ./... | grep /cmd/)

default:
	go build $(CMDS)

test:
	go clean $(PKGS)
	go test $(PKGS) -check.v -coverprofile=coverage.txt -covermode=atomic

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
	golangci-lint run