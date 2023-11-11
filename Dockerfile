FROM golang:1.20.7-alpine3.17

RUN apk update
RUN apk add make
RUN apk add openssl

WORKDIR /usr/local/app/src

COPY . .

# build
RUN make
RUN mv project-mirror ../
RUN cp -r public/ ../

WORKDIR /usr/local/app

# Remove build source
RUN rm -rf ./src

ENTRYPOINT [ "/usr/local/app/project-mirror" ]
