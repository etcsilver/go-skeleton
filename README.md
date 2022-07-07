# Estructura Projecto GO

*** Config ***
go mod edit -module github.com/etcsilver/go-name

*** Dependencias ***
```
go get -u github.com/gin-gonic/gin
go get -u github.com/rs/zerolog/log
go get -u github.com/joho/godotenv
go get -u gopkg.in/mgo.v2
go get -u github.com/zsais/go-gin-prometheus
go get -u github.com/swaggo/swag/cmd/swag
```

**Curl:**
```
### Pruebas Locales
curl -v -X POST \
  http://localhost:8081/hello \
  -H 'content-type: application/json' \
  -d '{ "name": "Fernando" }'

```

*** Swagger Doc ***
```
export PATH=$(go env GOPATH)/bin:$PAT
go get -u github.com/swaggo/swag/cmd/swag

swag init -g cmd/main.go -o /Pathers/../doc --ot json --instanceName apiname
```