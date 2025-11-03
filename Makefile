build:
	echo "Hello World!"

format:
	cd frontend && npm run format && npm run lint -- --fix

frontend-dev:
	cd frontend && npm run dev -- --open

backend-dev:
	cd backend/cmd && go run api/main.go