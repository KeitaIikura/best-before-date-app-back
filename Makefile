dc = docker-compose

build_dev:
	$(dc) up -d db
	make migrate
	make sqlboiler
	$(dc) build app

run:
	$(dc) up app

stopdb:
	$(dc) stop db

clean:
	$(dc) down -v --remove-orphans

migrate:
	$(dc) build migrate
	$(dc) run --rm migrate

sqlboiler:
	rm -rf ./pkg/dbmodels
	$(dc) build sqlboiler
	$(dc) run --rm sqlboiler /tmp/sqlboiler.sh
	cp -rp ./tmp/dbmodels/ ./pkg/dbmodels
	rm -rf ./tmp/dbmodels/

mock:
	go generate ./...

test:
	go test ./... -shuffle=on
