run-app-producer:
	@echo "Running app..."
	@go run cmd/producer/producer.go

run-app-consumer:
	@echo "Running app..."
	@go run cmd/consumer/consumer.go