run:
   	go run main.go
build:
   	docker build -t yoshikrit/jaeger-test . 
compose:
	docker-compose up -d