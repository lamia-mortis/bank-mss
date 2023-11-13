ifndef FRONTEND_SRC_DIR
	FRONTEND_SRC_DIR=resources
endif

ifndef SERVER_BINARY_FILE
	SERVER_BINARY_FILE=staticServer
endif

ifndef WORK_DIR
	WORK_DIR=/var/www/app
endif

ifndef NVM_VERSION
	NVM_VERSION="0.39.5"
endif

ifndef NODE_VERSION
	NODE_VERSION="18.16.1"
endif

ifndef UI_HOST
	UI_HOST=0.0.0.0
endif

SERVER_CMD_FILE=cmd/main.go

frontend_install:
	npm --prefix ${FRONTEND_SRC_DIR} install

frontend_run:
	npm run --prefix ${FRONTEND_SRC_DIR} serve

frontend_build:
	npm run --prefix ${FRONTEND_SRC_DIR} build

server_run:
	go mod tidy
	go run ${SERVER_CMD_FILE}

server_build:
	CGO_ENABLED=0 go build -o ${SERVER_BINARY_FILE} ${SERVER_CMD_FILE}

# build an image with a tag for pushing to the Docker Hub
ui_docker_image: # V=1.0
	docker build --build-arg="WORK_DIR=${WORK_DIR}" \
				 --build-arg="SERVER_BINARY_FILE=${SERVER_BINARY_FILE}" \
				 --build-arg="FRONTEND_SRC_DIR=${FRONTEND_SRC_DIR}" \
				 --build-arg="NVM_VERSION=${NVM_VERSION}" \
				 --build-arg="NODE_VERSION=${NODE_VERSION}" \
				 -t lamiamortis/ui:${V} .