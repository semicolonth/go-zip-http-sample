# Go sample code for generate zip and serve via HTTP
## Requirements
* Can be successfully run on go1.8.2 windows/amd64
## Configuration
Currently HTTP server ip address and port binding configuration is hard-coded as localhost:4001
## Run
```
$ go run zip.go
```
## Test
```
$ curl -vv http://localhost:4001/
```
or browse to that URL from web browser