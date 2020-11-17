# movie.api

##### Run api & mysql containers
```bash
docker-compose up --build -d
```

##### Run api container
```
MYSQL_DATABASE=homestead MYSQL_HOST=localhost MYSQL_USERNAME=root MYSQL_PASSWORD=root ./server
```