# WEB INTERFACE MICROSERVICE

### MANUAL SETUP

Directive should be run from the microservice root folder
- run static frontend on the `localhost:8080` (use `frontend_install` only if dependencies are absent locally):

```
    make frontend_install frontend_run
```
- run backend server on the `localhost:8080`, that serves static frontend build (use `frontend_install` only if dependencies are absent locally): 
```
    make frontend_install frontend_build server_run
```