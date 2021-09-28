test:
	go test -coverprofile projectCoverage.html ./...
build:
	env GOOS=linux go build  -o bin/main spotHeroProject
deploy:
	sls deploy
all: test build deploy

