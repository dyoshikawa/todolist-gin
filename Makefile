.PHONY: migrate create-migration build deploy

os=osx

migrate:
	./tool/migrate.${os} -source file://./db/migration \
		-database mysql://${user}:${password}@tcp\(${host}:${port}\)/${database} up

migrate-local:
	make migrate user=root password=secret host=0.0.0.0 database=dev port=3307
	make migrate user=root password=secret host=0.0.0.0 database=test port=3308

create-migration:
	./tool/migrate.${os} create -ext sql -dir ./db/migration ${name}

deploy: build
	heroku container:push web
	heroku container:release web
