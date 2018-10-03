DEPS := github.com/aws/aws-sdk-go/aws github.com/aws/aws-sdk-go/aws/awserr github.com/aws/aws-sdk-go/aws/session github.com/aws/aws-sdk-go/service/secretsmanager

build: dist/aws-sm-env

dist/aws-sm-env:
	mkdir -p dist
	go generate ./...
	go build -o dist/aws-sm-env ./cmd/aws-sm-env

deps: $(DEPS)

$(DEPS):
	go get $@

clean:
	rm -r dist
