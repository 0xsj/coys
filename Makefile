dev:
	docker-compose build && docker-compose up
down:
	docker-compose down --remove-orphans