run:
	cd backend && go run main.go

run-build:
	cd backend && ./companion

build:
	cd backend && go build && mv backend companion
	cd handler && npm run build

setup:
	cd backend && go get
	cd frontend && npm i
	cd handler && npm i
