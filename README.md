# server-spike

## Init Golang Package
```
go mod init server
export GOPROXY=https://goproxy.io
go mod tidy

dlv debug server/main.go
b path:line
```
## Guide
1. modify ./resources/demo.yaml
2. bin/gen_swagger_server.sh
3. source bin/configuration.sh
4. bin/run_gin_swagger.sh # go run tests/main.go
5. access swagger: http://localhost:9000/swagger/index.html



## Object hierarchy
- main.go: entry
### request flow
- api: Swagger groups
- service: HTTP path and function 
- domain: business codes related to service
- po: persitence object related to db
- db
```
api -- service -- domain -- po -- db
```
### functions
- init: load OS env and prepare
- pkg: extension methods
- remote: remote services
