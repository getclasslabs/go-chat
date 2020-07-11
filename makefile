all: init

init:
	sh scripts/init.sh

db:
	docker stack deploy -c build/mysql/docker-compose.yaml gct

jaegerDev:
	docker stack deploy -c build/jaeger/docker-compose.yaml gct

jaeger:
	docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest