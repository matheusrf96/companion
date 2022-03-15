run:
	cd backend && go run main.go

run-build:
	cd backend && ./go-webserver

build:
	cd backend && go build && mv backend go-webserver
	cd handler && npm run build

setup:
	cd backend && go get
	cd frontend && npm i
	cd handler && npm i
