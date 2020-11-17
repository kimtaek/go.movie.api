BINARY=engine
test:
	MYSQL_DATABASE=homestead MYSQL_HOST=host.docker.internal MYSQL_USERNAME=root MYSQL_PASSWORD=root go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} app/*.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t go-clean-arch .

docker-run:
	docker run -idt \
		-e MYSQL_DATABASE=homestead \
		-e MYSQL_HOST=host.docker.internal \
		-e MYSQL_USERNAME=root \
		-e MYSQL_PASSWORD=root \
		-p 8081:8080 \
		--name=youtha.movie.api.local \
		go-clean-arch

docker-rm:
	docker rm -f youtha.movie.api.local

run:
	docker-compose up --build -d

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint