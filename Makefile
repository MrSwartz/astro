build:
	docker-compose build astro

run:
	docker-compose up astro

first_migrate:
	migrate create -ext sql -dir ./schema -seq init  

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up

install_postgres:
	docker run --name=astro -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres   

# request examples
get_picday_date:
	curl --location --request GET '127.0.0.1:8000/picday?store=true&date=2021-02-02'

get_stored_range:
	curl --location --request GET 'http://127.0.0.1:8000/stored?start_date=2020-01-01&end_date=2024-02-11'

get_stored_date:
	curl --location --request GET 'http://127.0.0.1:8000/stored?date=2022-02-02'

scr:
	docker run --name=astro -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres  && 
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up &&
	