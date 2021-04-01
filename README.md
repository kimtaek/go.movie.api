# movie.api

##### Run api & mysql containers
```bash
docker-compose up --build -d
```

##### Run api container
```
MYSQL_DATABASE=homestead MYSQL_HOST=localhost MYSQL_USERNAME=root MYSQL_PASSWORD=root ./server
```

##### Install swag
```
go get -u github.com/swaggo/swag/cmd/swag
```

##### Goland swag auto generate
```
before lunch > external tool
    program: $GOPATH$/bin/swag
    arguments: init
    working dir: $ProjectFileDir$
```