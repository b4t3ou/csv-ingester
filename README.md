# csv-ingester

This is a small project where we have a storage application and a client communicating
through grpc

## Install

Clone the repository into the right folder
```
$GOPATH/src/github.com/b4t3ou/csv-ingester
```

Install dependencies

```
glide up
```

## Testing
### Unit testing

```bash
make test
```


### Integration testing

Run localstack first

```
docker-compose up localstack
```

Run all the tests

```
make test-all
```

## Running the applications

### Run the server (storage)

```
docker-compose up --build storage-server
```

### Run the file reader

```
docker-compose up --build storage-client
```



