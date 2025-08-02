tests:
	go test ./...

build:
	docker compose build m3u_consumer

run:
	docker compose run m3u_consumer
