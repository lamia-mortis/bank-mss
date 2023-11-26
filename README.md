# Mono Repository for all the microservices

## Set Up with Kubernetes
- `kubectl`, `make` and `minikube` should already be installed;
- make sure the `docker daemon`, `minikube` are running and `ingress`, `ingress-dns` addons are available;
```
    sudo systemctl start docker
    minikube start
    minikube addons enable ingress
    minikube addons enable ingress-dns
```
- `80` port of the host machine should not be in use; 
- add hosts to the `/etc/hosts`:
```
    127.0.0.1       ui
    127.0.0.1       gateway
```
- run the following command from the **project root**:
```
    make k8s_start
```
- the app will be available under the:
```
    http://ui
    http://gateway
```

## Set Up with Docker
- `docker engine`, `docker compose` and `make` should already be installed, `docker engine` should run;
- `8080` and `8888` ports of the host machine should not be in use; 
- create `.env` file using `.env.dist` as template;
- run the following command from the **project root**:
```
    make docker_start
```
- to stop containers run the following command from the **project root**:
```
    make docker_stop
```
- the app will be available under the following origins:
```
    http://localhost:8080 - ui
    http://localhost:8888 - gateway
```

## Set Up Manually
- `docker engine`, `docker compose` should already be installed, `docker engine` should run;
- required: 
    - `make 4.3`;
    - `nodejs 18.16.1`;
    - `protoc 3.21.12`;
    - `go 1.21.3`;
    - `golang-migrate 4.16.2`;
    - `sqlc 1.24.0`;
- `8080`, `8888`, `50051`, `9090`, `5432` ports of the host machine should not be in use; 
- during the initial setup (first run) execute the following command from the **project root**:
```
    make dev_env_setup
```
- to run the project with all microservices:
```
    make dev_env_up
```
- if it is the first run and the DB is empty: 
```
    make db_setup
```
- the app will be available under the following origins:
```
    http://localhost:8080  - ui
    http://localhost:8888  - gateway HTTP
    localhost:50051        - gateway gRPC
    localhost:9090         - auth gRPC
```