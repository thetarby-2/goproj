# How To Run With Docker
Db and api can be run with docker compose at the same time with;
```
$ docker-compose up
```
or 
```
$ docker compose up
```

# How To Run Tests
There are some tests under /app/tests folder. They can be run with;
```
$ go mod download // if dependencies are not installed yet.
$ cd app/tests
$ go test -v
```

