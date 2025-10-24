# Simple Makefile for Go backend + Svelte frontend

build:
	cd ./backend/cmd && go build -o server .

run:
	cd backend && go run .

frontend:
	cd frontend && npm run dev

frontend-build:
	cd frontend && npm run build
