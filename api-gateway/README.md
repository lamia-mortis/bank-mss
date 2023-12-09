# API-GATEWAY 

### MANUAL SETUP

Directive should be run from the microservice root folder
- create `.env` file, using `.env.dist` as a template;
- run microservice with *HTTP* server on the `localhost:8888` and *gRPC* server on the `localhost:50051`:
```
    make server_run
```