.PHONY: test run

run:
	docker-compose up -d --build

test:
	go test ./... -v

benchmark:
	go test ./... -bench=. -benchmem

down:
	docker-compose down