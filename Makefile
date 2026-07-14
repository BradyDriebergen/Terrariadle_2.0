build:
	cd frontend && npm run build
	rm -rf internal/web/build
	cp -r frontend/build internal/web/build
	go build -ldflags "-X main.version=$$(git describe --tags --always --dirty)" -o bin/terrariadle .

run:
	./bin/terrariadle