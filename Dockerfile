FROM golang:1.20.7-alpine3.17

ARG FILE_NAME

RUN apk update
RUN apk add make

WORKDIR /usr/local/app/src

COPY . .

# build
RUN make
RUN mv ${FILE_NAME} ../

WORKDIR /usr/local/app

# Remove build source
RUN rm -rf ./src

ENTRYPOINT [ "/usr/local/app/${FILE_NAME}" ]