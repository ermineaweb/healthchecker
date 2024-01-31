build:
	docker build -t healthchecker:latest .

serve:
	go run test/server.go