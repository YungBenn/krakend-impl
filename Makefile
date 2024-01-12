docker.up:
	docker-compose up -d

docker.down:
	docker-compose down

docker.clean:
	docker-compose kill && docker-compose rm -f
	docker rmi krakend-hello-service:v1
	docker rmi krakend-user-service:v1