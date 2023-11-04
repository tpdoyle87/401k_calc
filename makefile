up:
	docker compose up -d 401k_calculator_db
down:
	docker compose down 401k_calculator_db
build:
	docker compose build --no-cache
running:
	docker ps -a
app-up:
	docker compose up 401k-calculator-app
app-down:
	docker compose down 401k-calculator-app
stop:
	docker stop
delete:
	docker container prune -f