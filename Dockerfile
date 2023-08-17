FROM golang:1.20.7-alpine3.17

RUN apk update
RUN apk add make

WORKDIR /usr/local/app/src

COPY . .

# build
RUN make
RUN mv project-mirror ../
RUN cp -r static/ ../
RUN cp server.json ../

WORKDIR /usr/local/app

# Remove build source
RUN rm -rf ./src

ENTRYPOINT [ "/usr/local/app/project-mirror" ]