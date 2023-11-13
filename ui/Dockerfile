# base go image
FROM golang:1.21.3 as builder

# installing nvm with node and npm
ENV NVM_DIR /root/.nvm 
ARG NVM_VERSION=${NVM_VERSION}
ARG NODE_VERSION=${NODE_VERSION}

RUN mkdir -p ${NVM_DIR}
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v${NVM_VERSION}/install.sh | bash \
    && . ${NVM_DIR}/nvm.sh \
    && nvm install $NODE_VERSION \
    && nvm alias default $NODE_VERSION \
    && nvm use default

# add node and npm to path so the commands are available
ENV NODE_PATH $NVM_DIR/v${NODE_VERSION}/lib/node_modules
ENV PATH $NVM_DIR/versions/node/v${NODE_VERSION}/bin:${PATH}

ARG WORK_DIR=${WORK_DIR}

RUN mkdir -p ${WORK_DIR}

COPY . ${WORK_DIR}

WORKDIR ${WORK_DIR}

# build an app with static resources
ARG SERVER_BINARY_FILE=${SERVER_BINARY_FILE}
ARG FRONTEND_SRC_DIR=${FRONTEND_SRC_DIR}

RUN make frontend_install frontend_build server_build
RUN chmod +x ${WORK_DIR}/${SERVER_BINARY_FILE}

# build a tiny docker image
FROM alpine:3.18.4

ARG WORK_DIR=${WORK_DIR}

RUN mkdir -p ${WORK_DIR}

ARG SERVER_BINARY_FILE=${SERVER_BINARY_FILE}
ARG FRONTEND_SRC_DIR=${FRONTEND_SRC_DIR}

# using tmp dir to prevent volume owerwriting
COPY --from=builder ${WORK_DIR}/${SERVER_BINARY_FILE} ${WORK_DIR}
COPY --from=builder ${WORK_DIR}/${FRONTEND_SRC_DIR}/dist ${WORK_DIR}/${FRONTEND_SRC_DIR}/dist

ENV UI_HOST=${UI_HOST}

CMD [ "/var/www/app/staticServer" ]