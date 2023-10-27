migrate:
	docker-compose run app migrate -database postgres://${APOD_DB_USER}:${APOD_DB_PASSWORD}@db:${APOD_DB_PORT}/${APOD_DB_NAME}?sslmode=disable -path database/migration up

run:
	docker-compose build
	docker-compose up
	migrate