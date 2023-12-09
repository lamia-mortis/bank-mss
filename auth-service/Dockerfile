# base go image
FROM golang:1.21.3 as builder

ARG WORK_DIR=${WORK_DIR}
RUN mkdir -p ${WORK_DIR}
COPY . ${WORK_DIR}

WORKDIR ${WORK_DIR}

ARG AUTH_BINARY_FILE=${AUTH_BINARY_FILE}
RUN make server_build
RUN chmod +x ${WORK_DIR}/${AUTH_BINARY_FILE}

RUN apt install curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# build a tiny docker image
FROM alpine:3.18.4

ARG WORK_DIR=${WORK_DIR}
ARG AUTH_BINARY_FILE=${AUTH_BINARY_FILE}

WORKDIR ${WORK_DIR}

COPY --from=builder ${WORK_DIR}/${AUTH_BINARY_FILE} ${WORK_DIR}
COPY --from=builder ${WORK_DIR}/migrate.linux-amd64 ${WORK_DIR}/migrate

COPY .env.dist ${WORK_DIR}
RUN cp ${WORK_DIR}/.env.dist ${WORK_DIR}/.env

COPY wait-for.sh ${WORK_DIR}
RUN chmod 755 ${WORK_DIR}/wait-for.sh

COPY db/migrations ${WORK_DIR}/migrations

COPY entrypoint.sh ${WORK_DIR}
RUN chmod 755 ${WORK_DIR}/entrypoint.sh

# pass parameter to the entrypoint
CMD [ "/var/www/app/authApp" ] 
ENTRYPOINT [ "sh", "/var/www/app/entrypoint.sh" ]