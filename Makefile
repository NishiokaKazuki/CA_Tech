run-server:
	realize start --server

docker-build:
	cd docker &&\
	docker-compose build --no-cache

docker-up:
	cd docker &&\
	docker-compose up -d

docker-stop:
	cd docker &&\
	docker-compose stop

docker-rm:
	cd docker &&\
	docker-compose rm -f
	rm -rf db/data/*

docker-ps:
	cd docker &&\
	docker-compose ps

docker-exec-db:
	cd docker &&\
	docker exec -it ca_tech_db /bin/bash
