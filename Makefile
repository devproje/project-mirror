TARGET=project-mirror
SOURCES= ./src/*.go ./src/**/*.go
MAIN=./src/main.go

TAG=default

$(TARGET): $(MAIN) $(SOURCES)
	go build -o $(TARGET) $(MAIN)

debug:
	go run $(MAIN) -debug

docker:
	docker-compose build --no-cache --build-arg	FILE_NAME="$(TARGET)"

clean:
	rm $(TARGET)

rmlog:
	rm log.txt