# Multiplica

### Pre requites
* go version >= 1.5
* Docker
* curl

### Run app

Build and run the app

```
  docker build -t docker_multiplica .  
  docker run --name docker_multiplica -d -p 8080:8080 -p 9090:9090 docker_multiplica
```

You can call the endpoint using the following command:

```
curl -X GET http://localhost:8080/api/v1/multiply/12/4
```

and for grpc

```
grpcurl -d '{"numberA": 2, "numberB": 12}' -plaintext localhost:9090 grpc.MultiplyService/Multiply
```

### Tests

```
go mod download
go test ./...  
```

output

```
?   	multiplica/cmd	[no test files]
?   	multiplica/cmd/config	[no test files]
ok  	multiplica/internal/entity	(cached) [no tests to run]
ok  	multiplica/internal/handler/grpc	(cached)
ok  	multiplica/internal/handler/http	(cached)
ok  	multiplica/internal/service	(cached)
```

### Stress testing

Install the server locally
```
docker pull loadimpact/k6
```

Run stress tests 
```
export HOST_SERVER=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' docker_multiplica)
cd k6_tests
docker run -i loadimpact/k6 -e HOST_SERVER=$HOST_SERVER run - <script.js --out json=test.json
```
