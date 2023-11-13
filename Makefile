docker_start:
	docker compose up -d

docker_stop:
	docker compose down

k8s_start:
	kubectl apply -f k8s/microservices
	kubectl apply -f k8s/ingress.yml
	minikube tunnel

## dev environment
dev_env_setup:
	cd ui && make frontend_install && go mod tidy
	cd api-gateway && go mod tidy

dev_env_up: network postgres
	make -j2 dev_ui_up dev_gateway_up

dev_ui_up:
    ## using $() to make the jobserver available
	cd ui && $(MAKE) frontend_run && $(MAKE) server_run

dev_gateway_up:
	cd api-gateway && $(MAKE) server_run

network:
    ## check if the network already exists
	docker network inspect bank-mss >/dev/null 2>&1 || docker network create bank-mss

postgres:
	docker run --name postgres --network bank-mss -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
