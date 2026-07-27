build:
	cd frontend && npm run build
	rm -rf internal/web/build
	cp -r frontend/build internal/web/build
	go build -ldflags "-X main.version=$$(git describe --tags --always --dirty)" -o bin/terrariadle-test .

run:
	./bin/terrariadle-test

build-prod:
	cd frontend && npm run build
	rm -rf internal/web/build
	cp -r frontend/build internal/web/build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-trimpath \
		-ldflags "-s -w -X main.version=$$(git describe --tags --abbrev=0)" \
		-o bin/terrariadle .
