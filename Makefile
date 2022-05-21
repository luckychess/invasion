# build invasion simulator
build:
	go build -o invasion

# run tests
test: build
	go test ./...

# test run with 2 aliens and a small map
run-simple: build
	./invasion 2 sample/input.txt

# test run with 300 aliens and a big map
run-big: build
	./invasion 300 sample/input_big.txt

# remove built binary
clean:
	rm invasion

# build docker image
.PHONY: docker
docker:
	docker build -f docker/Dockerfile -t invasion .

# launch cli inside docker container
docker-run:
	docker run -it invasion
