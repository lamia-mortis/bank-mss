# AUTH MICROSERVICE

### MANUAL SETUP

Directive should be run from the microservice root folder
- create `.env` file, using `.env.dist` as a template;
- run *gRPC* server on the `localhost:9090`:
```
    make server_run
```
- if it is the first run and the `auth` DB does not exist (`postgres` container should running and `golang-migrate 4.16.2` installed): 
```
    make create_db migrate_up
```