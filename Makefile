build:
	cd frontend && npm run build
	cp -r frontend/build internal/web/build
	go build -o bin/terrariadle .

run:
	./bin/terrariadle